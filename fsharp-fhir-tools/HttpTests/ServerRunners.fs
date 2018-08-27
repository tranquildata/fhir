module ServerRunners

open System
open System.IO
open System.Threading
open System.Diagnostics
open Hl7.Fhir.Rest


type IFhirRunner =
    inherit IDisposable
    abstract member FhirClient: unit -> FhirClient
    abstract member Name: unit -> string


/// Static class giving access to the FHIR server currently being tested
type FhirServer private() =
    static member val Current : IFhirRunner = Unchecked.defaultof<IFhirRunner> with get, set


let private waitForFhirServer (runner: IFhirRunner) =
    eprintf "Waiting for %s to start" (runner.Name())
    let mutable brk = false
    while brk = false do
        try
            let cs  = runner.FhirClient().CapabilityStatement
            let srch  = runner.FhirClient().Search("Patient")
            brk <- true
            eprintf " --> started"
        with _ ->
            Thread.Sleep(1000)
            eprintf "."
    eprintfn ""

let private startProcess cmd args =
    let argString = String.concat " " args
    printfn "Starting %s %s" cmd argString
    let proc = Process.Start(ProcessStartInfo(cmd, argString))
    proc
            

type GoFhirMongoRunner() as this =
    // Drop fhir-bench mongo database
    let dropDatabaseProc = Process.Start(ProcessStartInfo("""c:\windows\system32\bash.exe""", """-c "mongo fhir-bench --eval db.dropDatabase\(\)" """))
    do dropDatabaseProc.WaitForExit()

    // Run gofhir
    //let path = """C:\Users\eug\go\src\github.com\eug48\fhir\fhir-server\fhir-server.exe"""
    let path = Path.Combine("..", "..", "..", "..", "fhir-server", "fhir-server.exe")
    let proc = startProcess path ["-port 5001 -mongodbHostPort localhost:27017 -enableXML -databaseName fhir-bench"]
    let fhirBase = "http://localhost:5002"
    //let fhirBase = "http://localhost:3002"
    let fhirClient() =
        new FhirClient(fhirBase,
                            verifyFhirVersion = false,
                            PreferredFormat = ResourceFormat.Json,
                            Timeout = 30 * 1000)
    do waitForFhirServer this

    interface IFhirRunner with
        member __.FhirClient() = fhirClient()
        member __.Name() = "GoFHIR"

    interface IDisposable with
        member __.Dispose() =
            proc.Kill()
            proc.WaitForExit()


type HapiDerbyRunner() as this =
    // delete HAPI data directory
    let dirPath = """C:\Users\eug\Downloads\hapi-fhir-3.4.0-cli"""
    let dataDir = Path.Combine(dirPath, "target")
    do if Directory.Exists(dataDir) then Directory.Delete(dataDir, recursive = true)

    // run HAPI
    let hapiCmdPath = Path.Combine(dirPath, "hapi-fhir-cli.cmd")
    let hapiArgs = [
        "run-server"
        "--reuse-search-results-milliseconds off"
        "-v dstu3"
        "--disable-referential-integrity"
        "-p 6001"
    ]
    let proc = startProcess hapiCmdPath hapiArgs
    let fhirBase = "http://localhost:6002/baseDstu3/"
    let fhirClient() = new FhirClient(fhirBase, verifyFhirVersion = false, PreferredFormat = ResourceFormat.Json)
    do waitForFhirServer this

    interface IFhirRunner with
        member __.FhirClient() = fhirClient()
        member __.Name() = "HAPI"

    interface IDisposable with
        member __.Dispose() =
            proc.Kill()
            proc.WaitForExit()




type PyroRunner() as this =
    // delete HAPI data directory
    let dirPath = """C:\Users\eug\Downloads\hapi-fhir-3.4.0-cli"""
    let dataDir = Path.Combine(dirPath, "target")
    //do if Directory.Exists(dataDir) then Directory.Delete(dataDir, recursive = true)

    // run HAPI
    let hapiCmdPath = Path.Combine(dirPath, "hapi-fhir-cli.cmd")
    let hapiArgs = [
        "run-server"
        "--reuse-search-results-milliseconds off"
        "-v dstu3"
        "--disable-referential-integrity"
        "-p 6001"
    ]
    //let proc = Process.Start(ProcessStartInfo(hapiCmdPath, String.concat " " hapiArgs))
    let fhirBase = "http://localhost:8888/fhir/"
    let fhirClient() = new FhirClient(fhirBase, verifyFhirVersion = false, PreferredFormat = ResourceFormat.Json)
    do waitForFhirServer this

    interface IFhirRunner with
        member __.FhirClient() = fhirClient()
        member __.Name() = "Pyro"

    interface IDisposable with
        member __.Dispose() =
            ()
            //proc.Kill()
            //proc.WaitForExit()



