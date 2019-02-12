package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
	"time"

	"github.com/eug48/fhir/auth"
	"github.com/eug48/fhir/fhir-server/middleware"
	"github.com/eug48/fhir/server"
)

var gitCommit string

func main() {
	port := flag.Int("port", 3001, "Port to listen on")
	reqLog := flag.Bool("reqlog", false, "Enables request logging -- use with caution in production")
	mongodbURI := flag.String("mongodbURI", "mongodb://mongo:27017/?replicaSet=rs0", "MongoDB connection URI - a replica set is required for transactions support")
	databaseName := flag.String("databaseName", "fhir", "MongoDB database name to use by default")
	enableMultiDB := flag.Bool("enableMultiDB", false, "Allow request to specify a specific Mongo database instead of the default, e.g. http://fhir-server/db/test4_fhir/Patient?name=alex")
	databaseSuffix := flag.String("databaseSuffix", "", "Request-specific MongoDB database name has to end with this (optional, e.g. '_fhir')")
	disableSearchTotals := flag.Bool("disableSearchTotals", false, "Don't query for all results of a search to return Bundle.total, only do paging")
	enableXML := flag.Bool("enableXML", false, "Enable support for the FHIR XML encoding")
	validatorURL := flag.String("validatorURL", "", "A FHIR validation endpoint to proxy validation requests to")
	failedRequestsDir := flag.String("failedRequestsDir", "", "Directory where to dump failed requests (e.g. with malformed json)")
	startMongod := flag.Bool("startMongod", false, "Run mongod (for 'getting started' docker images - development only)")
	flag.Parse()

	if *startMongod {
		// this is for the fhir-server-with-mongo docker image
		mongod := exec.Command("mongod", "--replSet", "rs0")
		err := mongod.Start()
		if err != nil {
			fmt.Fprintf(os.Stderr, "[server.go] ERROR: failed to start mongod: %#v\n", err)
			os.Exit(1)
		}
		go func() {
			err := mongod.Wait()
			fmt.Fprintf(os.Stdout, "[server.go] mongod has exited with %#v, also exiting\n", err)
			if err == nil {
				os.Exit(0)
			} else {
				os.Exit(1)
			}
		}()

		// wait for MongoDB
		fmt.Println("Waiting for MongoDB")
		for {
			conn, err := net.Dial("tcp", "127.0.0.1:27017")
			if err == nil {
				conn.Close()
				break
			} else {
				time.Sleep(1 * time.Second)
				fmt.Print(".")
			}
		}
		fmt.Println()

		// initiate the replica set
		time.Sleep(2 * time.Second)
		mongoShell := exec.Command("mongo", "--eval", "rs.initiate()")
		err = mongoShell.Start()
		if err != nil {
			fmt.Fprintf(os.Stderr, "[server.go] ERROR: failed to start mongo shell: %#v", err)
			os.Exit(1)
		}
		time.Sleep(2 * time.Second)
	}

	if *enableXML == false {
		fmt.Println("XML support is disabled (use --enableXML to enable)")
	}

	if gitCommit != "" {
		fmt.Printf("GoFHIR version %s\n", gitCommit)
	}
	fmt.Printf("MongoDB URI is %s\n", *mongodbURI)

	var MyConfig = server.Config{
		ServerURL:             fmt.Sprintf("http://localhost:%d", *port),
		IndexConfigPath:       "config/indexes.conf",
		DatabaseURI:           *mongodbURI,
		DefaultDatabaseName:   *databaseName,
		EnableMultiDB:         *enableMultiDB,
		DatabaseSuffix:        *databaseSuffix,
		DatabaseSocketTimeout: 2 * time.Minute,
		DatabaseOpTimeout:     90 * time.Second,
		DatabaseKillOpPeriod:  10 * time.Second,
		Auth:                  auth.None(),
		EnableCISearches:      true,
		CountTotalResults:     *disableSearchTotals == false,
		ReadOnly:              false,
		EnableXML:             *enableXML,
		EnableHistory:         true,
		Debug:                 true,
		ValidatorURL:          *validatorURL,
		FailedRequestsDir:     *failedRequestsDir,
	}
	s := server.NewServer(MyConfig)
	if *reqLog {
		s.Engine.Use(server.RequestLoggerHandler)
	}

	// Mutex middleware to work around the lack of proper transactions in MongoDB
	// (unless using a MongoDB >= 4.0 replica set)
	s.Engine.Use(middleware.ClientSpecifiedMutexesMiddleware())

	s.Run()
}
