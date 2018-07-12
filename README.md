Intervention Engine FHIR Server [![Build Status](https://travis-ci.org/intervention-engine/fhir.svg?branch=stu3_mar2017)](https://travis-ci.org/intervention-engine/fhir) [![GoDoc](https://godoc.org/github.com/intervention-engine/fhir?status.svg)](https://godoc.org/github.com/intervention-engine/fhir)
===================================================================================================================================================================

This project provides [HL7 FHIR STU3](http://hl7.org/fhir/STU3/index.html) models and server components implemented in Go and using MongoDB as storage. **This branch is the official STU3 release**. This is a
library that can be embedded into other server applications. The library is not a complete implementation of FHIR, as features that are selected are driven by the
[Intervention Engine](https://github.com/intervention-engine/ie), [eCQM Engine](https://github.com/mitre/ecqm), [Patient Matching Test Harness](https://github.com/mitre/ptmatch)
and [Synthetic Mass](https://github.com/synthetichealth/syntheticmass) projects.

Currently, this server library supports:

-	JSON representations of all resources
-	XML representations of all resources via [FHIR.js](https://github.com/lantanagroup/FHIR.js) (except for primitive extension)
-	Create/Read/Update/Delete (CRUD) operations
-	Conditional update and delete
-	Resource-level versioning and history operations
-	Arbitrary-precision storage for decimals
-	Some but not all search features
	-	All defined resource-specific search parameters except composite types and contact (email/phone) searches
	-	Chained searches
	-	Reverse chained searches using `_has`
	-	`_include` and `_revinclude` searches (*without* `_recurse`)
-	Batch bundle uploads (POST, PUT, and DELETE entries)

Currently, this server does *not* support the following major features:

-	Transactional isolation (see workaround below)
-	XML representation of primitive extensions


Transactions
-----------

MongoDB is used as the underlying database and has recently acquired multi-document transaction features in MongoDB 4.0. This can be supported by this project once implemented in the new offially-supported Go driver - see https://github.com/mongodb/mongo-go-driver.

In the meantime the accompanying [fhir-server](https://github.com/eug48/fhir-server) project implements a workaround. Clients can send a `X-Mutex-Name` header and two requests with the same value of this header should execute serially.


Development
-----------

This project uses Go 1.10. To test the library, first, install all of the dependencies:

```
$ go get -t ./...
```

Once the dependecies have been installed, you can invoke the test suite by running:

```
$ go test ./...
```

Usage
-----

Users of this library should work with the [FHIRServer](https://godoc.org/github.com/intervention-engine/fhir/server#FHIRServer) struct. Web request
handlers in this library are implemented using [Gin](https://gin-gonic.github.io/gin/).

Examples of usage can be found in the [server set up of the eCQM Engine](https://github.com/mitre/ecqm/blob/master/server.go), the
[server set up of Intervention Engine](https://github.com/intervention-engine/ie/blob/master/server.go), or the [GoFHIR server used by SyntheticMass](https://github.com/synthetichealth/gofhir/blob/master/main.go).

License
-------

Copyright 2017 The MITRE Corporation

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
