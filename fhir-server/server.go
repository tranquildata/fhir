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
	mongodbURI := flag.String("mongodbURI", "mongodb://localhost:27017/?replicaSet=rs0", "MongoDB connection URI - a replica set is required for transactions support")
	startMongod := flag.Bool("startMongod", false, "Run mongod (for 'getting started' docker images - development only)")
	databaseName := flag.String("databaseName", "fhir", "MongoDB database name to use")
	enableXML := flag.Bool("enableXML", false, "Enable support for the FHIR XML encoding")
	validatorURL := flag.String("validatorURL", "", "A FHIR validation endpoint to proxy validation requests to")
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
		DatabaseName:          *databaseName,
		DatabaseSocketTimeout: 2 * time.Minute,
		DatabaseOpTimeout:     90 * time.Second,
		DatabaseKillOpPeriod:  10 * time.Second,
		Auth:                  auth.None(),
		EnableCISearches:      true,
		CountTotalResults:     true,
		ReadOnly:              false,
		EnableXML:             *enableXML,
		EnableHistory:         true,
		Debug:                 true,
		ValidatorURL:          *validatorURL,
	}
	s := server.NewServer(MyConfig)
	if *reqLog {
		s.Engine.Use(server.RequestLoggerHandler)
	}

	// Mutex middleware to work around the lack of proper transactions in MongoDB (at least until MongoDB 4.0)
	s.Engine.Use(middleware.ClientSpecifiedMutexesMiddleware())

	s.Run()
}
