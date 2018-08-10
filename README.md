GoFHIR - Intervention Engine FHIR Server
=====================================================================

This project provides [HL7 FHIR STU3](http://hl7.org/fhir/STU3/index.html) models and server components implemented in Go and using MongoDB for storage. It aims to work well when multiple instances are run in containers - starting quickly, having modest resource utilisation and giving good performance.

This is a continuation of the [FHIR Server](https://github.com/mitre/fhir-server) developed by the MITRE Corporation. It is not a complete implementation of FHIR, as features are driven by the
[Intervention Engine](https://github.com/intervention-engine/ie), [eCQM Engine](https://github.com/mitre/ecqm), [Patient Matching Test Harness](https://github.com/mitre/ptmatch),
[Synthetic Mass](https://github.com/synthetichealth/syntheticmass) and [Patient Assistance Tool](https://www.patsoftware.com.au/) projects.

Currently this server should be considered experimental, with preliminary support for:

-	JSON representations of all resources
-	XML representations of all resources via [FHIR.js](https://github.com/lantanagroup/FHIR.js) (except for primitive extensions)
-	Create/Read/Update/Delete (CRUD) operations with versioning
-	Conditional update and delete
-	Resource-level history (basic support - lacks paging and filtering)
-	Batch bundles (POST, PUT and DELETE entries)
-	Arbitrary-precision storage for decimals
-	Some search features
	-	All defined resource-specific search parameters except composite types and contact (email/phone) searches
	-	Chained searches
	-	Reverse chained searches using `_has`
	-	`_include` and `_revinclude` searches (*without* `_recurse`)

Currently this server does *not* support the following major features:

-	Transactional isolation (see workaround below)
-	Validation (work in progress)
-	Resource summaries
-	Advanced search features
	-	Custom search parameters
	-	Full-text search
	-	Filter expressions
-	GraphQL

The following relatively basic items are next in line for development:

- Conditional reads (`If-Modified-Since` and `If-None-Match`)
- History support for paging, `_since`, `_at` and `_count`
- Batch interdependency validation
- Search for quantities with the system unspecified (i.e. by both unit and code)

Users are strongly encouraged to test thoroughly and contributions (including more tests) would be most welcome.


Transactions
--------------

MongoDB is used as the underlying database and has recently acquired multi-document transaction features in version 4.0. This can be supported by this project once implemented in the new officially-supported Go driver - see https://github.com/mongodb/mongo-go-driver.

Another approach would be to create an alternative backend, perhaps using PostgreSQL, ArangoDB or Dgraph.

In the meantime this project implements a partial workaround. Clients can send a `X-Mutex-Name` header and two requests with the same value of this header will execute serially (provided there is only one active instance of this server). Please note that this won't give you the all-or-nothing behaviour of real transactions.

Getting started using Docker
-------------------------------

1. Install Docker
2. Run an image of this FHIR server that includes MongoDB, deleting all data after exiting:
		
		docker run --rm -it -p 3001:3001 gcr.io/eug48-fhir/fhir-server-with-mongo


3. You can also run MongoDB and this FHIR server in separate containers:

		docker run --name fhir-mongo -v /my/own/datadir:/data/db -d mongo
		docker run --rm -it --link fhir-mongo:mongo -e GIN_MODE=release -e MONGO_HOSTPORT=fhir-mongo:27017 -p 3001:3001 gcr.io/eug48-fhir/fhir-server


See MongoDB's Docker image documentation for more information, including how to persist data: https://hub.docker.com/_/mongo/


Building and running from source
---------------------------------

1. Install the Go programming language (at least version 1.10)
2. Install and start MongoDB
3. Install the [Dep](https://github.com/golang/dep) package-management tool for Go: https://golang.github.io/dep/docs/installation.html
4. Clone or download this repository under GOPATH/src (run `go env` to see your current GOPATH)
5. Run `dep ensure -vendor-only -v` to download dependencies into a `vendor` sub-directory (can take a long time)
6. Build and run the `fhir-server` executable

		$ cd fhir-server
		$ go build
		$ ./fhir-server --help
		Usage of ./fhir-server:
		-databaseName string
				MongoDB database name to use (default "fhir")
		-enableXML
				Enable support for the FHIR XML encoding
		-mongodbHostPort string
				MongoDB host:port (default "localhost:27017")
		-port int
				Port to listen on (default 3001)
		-reqlog
				Enables request logging -- use with caution in production
		-startMongod
				Run mongod (for 'getting started' docker images - development only)



If you wish to test the server with synthetic patient data, please reference [Generate and Upload Synthetic Patient Data](https://github.com/intervention-engine/ie/blob/master/docs/dev_install.md#generate-and-upload-synthetic-patient-data).


Building docker images
--------------------------

Dockerfiles as well as a Google Cloud Container Builder `cloudBuild.json` spec are in the root of the repository.

To build from the Dockerfiles:

```
$ docker build -t my-fhir-server:2018-07-12a .
$ docker build -f Dockerfile-with-mongo -t my-fhir-server-with-mongo:2018-07-12a .
```

To build using Google Container Builder (after setting up the gcloud tool):

```
# The main Dockerfile
$ gcloud container builds submit -t gcr.io/your-project-id/fhir-server:2018-07-12a

# Both Dockerfiles
$ gcloud container builds submit --config=cloudBuild.json  --substitutions=COMMIT_SHA=2018-07-12a
```

You can also set up a trigger via https://console.cloud.google.com/gcr/triggers 


Development
-------------

This project uses Go 1.10. To test the library first install dependencies by running `dep ensure -vendor-only  -v`. You can also try using the older `go get -t ./...` (this requires the project to be at the right place in your GOPATH).

Once the dependecies have been installed, you can invoke the test suite by running:

```
$ go test ./...
```

To run the server whilst it is under development it is often easier to combine the build and run steps into a single command from the `fhir-server` directory:

```
$ go run server.go
```

As a library
--------------

This package can also be used as a library. Examples of usage can be found in the [server set up of the eCQM Engine](https://github.com/mitre/ecqm/blob/master/server.go), the
[server set up of Intervention Engine](https://github.com/intervention-engine/ie/blob/master/server.go), or the [GoFHIR server used by SyntheticMass](https://github.com/synthetichealth/gofhir/blob/master/main.go).

License
-------

Copyright 2017 The MITRE Corporation

Copyright 2018 PAT Pty Ltd

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
