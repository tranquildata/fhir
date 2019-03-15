package server

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/eug48/fhir/models2"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/pkg/errors"
)

type AfterRoutes func(*gin.Engine)

type FHIRServer struct {
	Config           Config
	Engine           *gin.Engine
	MiddlewareConfig map[string][]gin.HandlerFunc
	AfterRoutes      []AfterRoutes
	Interceptors     map[string]InterceptorList
}

func (f *FHIRServer) AddMiddleware(key string, middleware gin.HandlerFunc) {
	f.MiddlewareConfig[key] = append(f.MiddlewareConfig[key], middleware)
}

// AddInterceptor adds a new interceptor for a particular database operation and FHIR resource.
// For example:
// AddInterceptor("Create", "Patient", patientInterceptorHandler) would register the
// patientInterceptorHandler methods to be run against a Patient resource when it is created.
//
// To run a handler against ALL resources pass "*" as the resourceType.
//
// Supported database operations are: "Create", "Update", "Delete"
func (f *FHIRServer) AddInterceptor(op, resourceType string, handler InterceptorHandler) error {

	if op == "Create" || op == "Update" || op == "Delete" {
		f.Interceptors[op] = append(f.Interceptors[op], Interceptor{ResourceType: resourceType, Handler: handler})
		return nil
	}
	return fmt.Errorf("AddInterceptor: unsupported database operation %s", op)
}

func NewServer(config Config) *FHIRServer {
	server := &FHIRServer{
		Config:           config,
		MiddlewareConfig: make(map[string][]gin.HandlerFunc),
		Interceptors:     make(map[string]InterceptorList),
	}
	server.Engine = gin.Default()

	if config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.DisableConsoleColor()

	server.Engine.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type, If-Match, If-None-Exist",
		ExposedHeaders:  "Location, ETag, Last-Modified",
		MaxAge:          86400 * time.Second, // Preflight expires after 1 day
		Credentials:     true,
		ValidateHeaders: false,
	}))

	if config.EnableXML {
		server.Engine.Use(EnableXmlToJsonConversionMiddleware())
		server.Engine.Use(AbortNonFhirXMLorJSONRequestsMiddleware)
	} else {
		server.Engine.Use(AbortNonJSONRequestsMiddleware)
	}

	if config.ReadOnly {
		server.Engine.Use(ReadOnlyMiddleware)
	}

	return server
}

func (f *FHIRServer) InitEngine() {
	var err error

	// Establish initial connection to mongo
	client, err := mongo.Connect(context.Background(), f.Config.DatabaseURI)
	if err != nil {
		panic(errors.Wrap(err, "connecting to MongoDB"))
	}

	getFCV := bson.NewDocument(
		bson.EC.Int32("getParameter", 1),
		bson.EC.Int32("featureCompatibilityVersion", 1),
	)
	fcvReader, err := client.Database("admin").RunCommand(context.TODO(), getFCV)
	if err != nil {
		fmt.Printf("MongoDB: unable to read featureCompatibilityVersion: %s\n", err.Error())
	} else {
		fcv, err := fcvReader.Lookup("featureCompatibilityVersion", "version")
		if err != nil {
			panic(errors.Wrap(err, "loading featureCompatibilityVersion"))
		}
		fmt.Printf("MongoDB: featureCompatibilityVersion %s\n", fcv.Value().StringValue())
	}

	log.Printf("MongoDB: Connected (default database %s)\n", f.Config.DefaultDatabaseName)

	// Pre-create collections for transactions
	db := client.Database(f.Config.DefaultDatabaseName)
	CreateCollections(db)

	// Ensure all indexes
	NewIndexer(f.Config.DefaultDatabaseName, f.Config).ConfigureIndexes(db)

	// Kick off the database op monitoring routine. This periodically checks db.currentOp() and
	// kills client-initiated operations exceeding the configurable timeout. Do this AFTER the index
	// build to ensure no index build processes are killed unintentionally.

	// ticker := time.NewTicker(f.Config.DatabaseKillOpPeriod)
	// TODO: disabled as requires high-grade permissions. Remove completely?
	// go killLongRunningOps(ticker, client.ConnectionString(), "admin", f.Config)

	// Register all API routes
	RegisterRoutes(f.Engine, f.MiddlewareConfig, NewMongoDataAccessLayer(client, f.Config.DefaultDatabaseName, f.Config.EnableMultiDB, f.Config.DatabaseSuffix, f.Interceptors, f.Config), f.Config)

	for _, ar := range f.AfterRoutes {
		ar(f.Engine)
	}

	// If not in -readonly mode, clear the count cache
	if !f.Config.ReadOnly {
		dbNames, err := client.ListDatabaseNames(context.TODO(), nil)
		if err != nil {
			panic(fmt.Sprint("Server: Failed to call ListDatabaseNames", err))
		}
		dbNames = append(dbNames, f.Config.DefaultDatabaseName)

		for _, databaseName := range dbNames {
			if strings.HasSuffix(databaseName, f.Config.DatabaseSuffix) {
				db := client.Database(databaseName)
				count, err := db.Collection("countcache").Count(context.Background(), nil)
				if count > 0 || err != nil {
					err = db.Collection("countcache").Drop(context.Background())
					if err != nil {
						panic(fmt.Sprintf("Server: Failed to clear count cache (%+v)", err))
					}
				}
			}
		}
	} else {
		log.Println("Server: Running in read-only mode")
	}
}

func (f *FHIRServer) Run() {
	f.InitEngine()

	url, err := url.Parse(f.Config.ServerURL)
	if err != nil {
		panic("Server: Failed to parse ServerURL: " + f.Config.ServerURL)
	}
	f.Engine.Run(":" + url.Port())
}

func (f *FHIRServer) InitDB(databaseName string) {
	// Connect
	client, err := mongo.Connect(context.Background(), f.Config.DatabaseURI)
	if err != nil {
		panic(errors.Wrap(err, "connecting to MongoDB"))
	}

	// Pre-create collections for transactions
	db := client.Database(databaseName)
	CreateCollections(db)

	// Ensure all indexes
	NewIndexer(databaseName, f.Config).ConfigureIndexes(db)
}

func CreateCollections(db *mongo.Database) {
	// MongoDB transactions require that collections be pre-created
	var err error
	for _, name := range models2.AllFhirResourceCollectionNames() {
		// fmt.Printf("pre-creating collection %s, %s\n", name, name+"_prev")
		_, err = db.RunCommand(context.Background(), bson.NewDocument(bson.EC.String("create", name+"_prev")))
		if err != nil && !strings.Contains(err.Error(), "already exists") {
			panic(err)
		}
		_, err = db.RunCommand(context.Background(), bson.NewDocument(bson.EC.String("create", name)))
		if err != nil && !strings.Contains(err.Error(), "already exists") {
			panic(err)
		}
	}
}
