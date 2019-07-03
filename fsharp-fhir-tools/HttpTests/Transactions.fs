module Transactions

open System
open Expecto
open Hl7.Fhir.Model
open Hl7.Fhir.Rest
open Utils

type OrganizationsFixture(fhir: FhirClient) =
    let autoDelete = new FhirAutoDelete(fhir)
    let create = autoDelete.Create
    interface IDisposable with member __.Dispose() = autoDelete.DeleteAll()
    
    member val O1 =
        Organization(
            Identifier = L [ Identifier("http://example.org/1", "1")],
            Active = N true
        ) |> create

    member val O2 =
        Organization(
            Identifier = L [ Identifier("http://example.org/1", "2")],
            Active = N true
        ) |> create

    member val O3 =
        Organization(
            Identifier = L [ Identifier("http://example.org/1", "3")],
            Active = N false
        ) |> create


let assertResourcesEqual (r1: #Resource) (r2: #Resource) =
    let jsonSerialiser = Hl7.Fhir.Serialization.FhirJsonSerializer()
    let j1 = jsonSerialiser.SerializeToString r1
    let j2 = jsonSerialiser.SerializeToString r2
    Expect.equal j1 j2 "assertResourcesEqual"

let transactionTests = [
    sprintf "transactions - GET/POST - ok",
        fun (fhir: FhirClient) () ->

            use organisations = new OrganizationsFixture(fhir)

            let bundle =
                Bundle(
                    Type = N Bundle.BundleType.Transaction,
                    Entry = L [
                        Bundle.EntryComponent(
                            Request = Bundle.RequestComponent(
                                Url =  "Patient/DoesntExist",
                                Method = N Bundle.HTTPVerb.GET
                            )
                        )
                        Bundle.EntryComponent(
                            Resource = Patient(
                                Identifier = L [ Identifier("http://example.com", "Patient1")],
                                ManagingOrganization = ResourceReference("Organization?active=false")
                            ),
                            Request = Bundle.RequestComponent(
                                Url =  "Patient",
                                Method = N Bundle.HTTPVerb.POST
                            )
                        )
                        Bundle.EntryComponent(
                            Resource = Patient(
                                Identifier = L [ Identifier("http://example.com", "Patient2")],
                                ManagingOrganization = ResourceReference("Organization/" + organisations.O3.Id)
                            ),
                            Request = Bundle.RequestComponent(
                                Url =  "Patient",
                                Method = N Bundle.HTTPVerb.POST
                            )
                        )
                    ]
                )
            
            let patientCount1 = fhir.Search("Patient").Total.Value
            let result = fhir.Transaction bundle
            let patientCount2 = fhir.Search("Patient").Total.Value
            Expect.equal (patientCount1 + 2) patientCount2 "transaction has added patients"
            Expect.equal result.Type (N Bundle.BundleType.TransactionResponse) "result bundle type"
            Expect.equal result.Entry.Count 3 "result bundle count"

            Expect.stringStarts result.Entry.[0].Response.Status "404" "0. 404"
            Expect.stringStarts result.Entry.[1].Response.Status "201" "1. 201"
            Expect.stringStarts result.Entry.[2].Response.Status "201" "2. 201"

            let patient1Location = result.Entry.[1].Response.Location
            let patient2Location = result.Entry.[2].Response.Location

            let patient1 = fhir.Get(patient1Location) :?> Patient
            let patient2 = fhir.Get(patient2Location) :?> Patient

            Expect.notEqual patient1.Id patient2.Id "patients different ids"
            Expect.equal patient1.ManagingOrganization.Reference patient2.ManagingOrganization.Reference "patients same ManagingOrganization"

            fhir.Delete(patient1)
            fhir.Delete(patient2)

    
    sprintf "transactions - GET/POST - conditional reference resolves to multiple",
        fun (fhir: FhirClient) () ->

            use organisations = new OrganizationsFixture(fhir)

            let bundle =
                Bundle(
                    Type = N Bundle.BundleType.Transaction,
                    Entry = L [
                        Bundle.EntryComponent(
                            Request = Bundle.RequestComponent(
                                Url =  "Patient/DoesntExist",
                                Method = N Bundle.HTTPVerb.GET
                            )
                        )
                        Bundle.EntryComponent(
                            Resource = Patient(
                                Identifier = L [ Identifier("http://example.com", "Patient1")],
                                ManagingOrganization = ResourceReference("Organization?active=true") // reference with multiple matches
                            ),
                            Request = Bundle.RequestComponent(
                                Url =  "Patient",
                                Method = N Bundle.HTTPVerb.POST
                            )
                        )
                        Bundle.EntryComponent(
                            Resource = Patient(
                                Identifier = L [ Identifier("http://example.com", "Patient2")],
                                ManagingOrganization = ResourceReference("Organization/" + organisations.O1.Id)
                            ),
                            Request = Bundle.RequestComponent(
                                Url =  "Patient",
                                Method = N Bundle.HTTPVerb.POST
                            )
                        )
                    ]
                )
            
            let patientCount1 = fhir.Search("Patient").Total.Value
            let exn = Expect.throwsC (fun () -> fhir.Transaction bundle |> ignore) (id)
            let patientCount2 = fhir.Search("Patient").Total.Value
            Expect.isNotNull exn "transaction fails"
            Expect.equal patientCount1 patientCount2 "transaction has made no changes"
            //Expect.equal result.Type (N Bundle.BundleType.TransactionResponse) "result bundle type"
            //Expect.equal result.Entry.Count 1 "result bundle count"

    sprintf "transactions - GET/POST - conditional create resolves to multiple",
        fun (fhir: FhirClient) () ->

            use organisations = new OrganizationsFixture(fhir)

            let bundle =
                Bundle(
                    Type = N Bundle.BundleType.Transaction,
                    Entry = L [
                        Bundle.EntryComponent(
                            Request = Bundle.RequestComponent(
                                Url =  "Patient/DoesntExist",
                                Method = N Bundle.HTTPVerb.GET
                            )
                        )
                        Bundle.EntryComponent(
                            Resource = Patient(
                                Identifier = L [ Identifier("http://example.com", "Patient1")]
                            ),
                            Request = Bundle.RequestComponent(
                                Url =  "Patient",
                                Method = N Bundle.HTTPVerb.POST
                            )
                        )
                        Bundle.EntryComponent(
                            Resource = Organization(
                                Identifier = L [ Identifier("http://example.org/1", "should-error-412")],
                                Active = N true
                            ),
                            Request = Bundle.RequestComponent(
                                Url =  "Organization",
                                Method = N Bundle.HTTPVerb.POST,
                                IfNoneExist = "active=true"
                            )
                        )
                    ]
                )
            
            let patientCount1 = fhir.Search("Patient").Total.Value
            let exn = Expect.throwsC (fun () -> fhir.Transaction bundle |> ignore) (id)
            let patientCount2 = fhir.Search("Patient").Total.Value
            Expect.isNotNull exn "transaction fails"
            Expect.equal patientCount1 patientCount2 "transaction has made no changes"
            //Expect.equal result.Type (N Bundle.BundleType.TransactionResponse) "result bundle type"
            //Expect.equal result.Entry.Count 1 "result bundle count"
    ]


[<Tests>]
let transactionsTest =
    let fhir = ServerRunners.FhirServer.Current.FhirClient()
    let tests = [
        testParam (fhir) (transactionTests |> List.toSeq )
    ]
    testList "Transactions" (List.collect id (tests |> List.map Seq.toList))
