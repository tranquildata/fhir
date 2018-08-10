open Expecto
open ServerRunners

[<EntryPoint>]
let main argv =
    System.Net.HttpWebRequest.DefaultWebProxy <- null
    Bogus.Randomizer.Seed <- System.Random(7777777)

    use runner = new GoFhirMongoRunner() :> IFhirRunner
    FhirServer.Current <- runner

    let expectoConfig = { defaultConfig with ``parallel`` = false; verbosity = Expecto.Logging.Verbose }

    let r = Tests.runTestsInAssembly expectoConfig argv
    printfn "done"
    r