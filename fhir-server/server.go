package main

import (
	"flag"
	"fmt"
	"time"
	"os"
	"os/exec"

	"github.com/eug48/fhir/fhir-server/middleware"
	"github.com/eug48/fhir/auth"
	"github.com/eug48/fhir/server"
)

var gitCommit string

func main() {
	port := flag.Int("port", 3001, "Port to listen on")
	reqLog := flag.Bool("reqlog", false, "Enables request logging -- use with caution in production")
	mongodbHostPort := flag.String("mongodbHostPort", "localhost:27017", "MongoDB host:port")
	startMongod := flag.Bool("startMongod", false, "Run mongod (for 'getting started' docker images - development only)")
	databaseName := flag.String("databaseName", "fhir", "MongoDB database name to use")
	enableXML := flag.Bool("enableXML", false, "Enable support for the FHIR XML encoding")
	validatorURL := flag.String("validatorURL", "", "A FHIR validation endpoint to proxy validation requests to")
	flag.Parse()

	if *startMongod {
		mongod := exec.Command("mongod")
		err := mongod.Start()
		if err != nil {
			fmt.Fprintf(os.Stderr, "[server.go] ERROR: failed to start mongod: %#v", err)
			os.Exit(1)
		}
		go func() {
			err := mongod.Wait()
			fmt.Fprintf(os.Stdout, "[server.go] mongod has exited with %#v, also exiting", err)
			if err == nil {
				os.Exit(0)
			} else {
				os.Exit(1)
			}
		}()
	}

	if *enableXML == false {
		fmt.Println("XML support is disabled (use --enableXML to enable)")
	}

	if gitCommit != "" {
		fmt.Printf("GoFHIR version %s\n", gitCommit)
	}
	fmt.Printf("MongoDB: host is %s\n", *mongodbHostPort)

	var MyConfig = server.Config{
		ServerURL:             fmt.Sprintf("http://localhost:%d", *port),
		IndexConfigPath:       "config/indexes.conf",
		DatabaseHost:          *mongodbHostPort,
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
