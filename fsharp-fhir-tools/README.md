Some FHIR projects developed in F# - enabling use of the FHIR .NET API library, JSON loading with automatic static typing and other goodness.

* HttpTests - integration tests and benchmarks for FHIR REST servers - emphasis on testing new GoFHIR features
* PathsByType - generates code for GoFHIR (fhir_types.go) from FHIR spec definitions, mapping paths to their types (e.g. Account.coverage.coverage --> Reference)