package runsvc

import (
	"strings"

	"github.com/andodeki/gen_project/builder/codegen_builder/src/data"
	"github.com/andodeki/gen_project/builder/codegen_builder/src/process"
)

func RunningServices(services []string, projectDefaults data.Data) {
	cts := make([]string, 0)
	addFiles := []string{"pprof"} //adding extra files not in configuration yaml file
	services = append(services, addFiles...)
	for _, v := range services {
		cts = append(cts, strings.ToLower(strings.Replace(v, "Clients", "", -1)))
	}
	// fmt.Printf("cts: %v\n", services)
	dataRunSvc := data.Data{
		ProductTargets: data.ProductTarget{
			OutputFilename: cts,
			DSource: data.Datasource{
				// DBSMap:       kvdbs,
				// DBSMapFields: dsrcs.DSource.DBSMapFields,
				DBS: services,
				// IMethods:     iterfacemethods,
			},
			// ProductFileTypeFrom: dataYAML.ProductFileType,
		},
		ProductSource:   projectDefaults.ProductSource,
		Company:         projectDefaults.Company,
		ProductFileType: ".go",
		FolderPathName:  projectDefaults.ProjectFolder + "/src/runsvc/",
		ConcreteTargets: cts,
	}
	// process.ProcessNonConcreteTargets("src/running.tmpl", dataRunSvc)
	process.ProcessConcreteTargets("src/runsvc/configuration.tmpl", dataRunSvc)
	process.ProcessConcreteTargets("src/runsvc/pprof.tmpl", dataRunSvc)
	process.ProcessConcreteTargets("src/runsvc/httpserver.tmpl", dataRunSvc)
	process.ProcessConcreteTargets("src/runsvc/pasetotoken.tmpl", dataRunSvc)
	process.ProcessConcreteTargets("src/runsvc/email.tmpl", dataRunSvc)
	process.ProcessConcreteTargets("src/runsvc/queue.tmpl", dataRunSvc)
	process.ProcessConcreteTargets("src/runsvc/postgresdb.tmpl", dataRunSvc)
	process.ProcessConcreteTargets("src/runsvc/scylladb.tmpl", dataRunSvc)
	process.ProcessConcreteTargets("src/runsvc/monetdb.tmpl", dataRunSvc)
	process.ProcessConcreteTargets("src/runsvc/redisdb.tmpl", dataRunSvc)
	process.ProcessConcreteTargets("src/runsvc/jaeger.tmpl", dataRunSvc)
	//[httpserver scylladb pasetotoken monetdb jaeger configuration postgresdb redisdb email queue]
	// for _, v := range dataRunSvc.ProductTargets.OutputFilename {
	// 	if v == "configuration" {
	// 		process.ProcessNonConcreteTargets("src/runsvc/configuration.tmpl", dataRunSvc, v)

	// 	} else if v == "httpserver" {
	// 		process.ProcessNonConcreteTargets("src/runsvc/httpserver.tmpl", dataRunSvc, v)

	// 	} else if v == "pasetotoken" {
	// 		process.ProcessNonConcreteTargets("src/runsvc/pasetotoken.tmpl", dataRunSvc, v)

	// 	} else if v == "email" {
	// 		process.ProcessNonConcreteTargets("src/runsvc/email.tmpl", dataRunSvc, v)

	// 	} else if v == "queue" {
	// 		process.ProcessNonConcreteTargets("src/runsvc/queue.tmpl", dataRunSvc, v)

	// 	} else if v == "postgresdb" {
	// 		process.ProcessNonConcreteTargets("src/runsvc/postgresdb.tmpl", dataRunSvc, v)

	// 	} else if v == "scylladb" {
	// 		process.ProcessNonConcreteTargets("src/runsvc/scylladb.tmpl", dataRunSvc, v)

	// 	} else if v == "monetdb" {
	// 		process.ProcessNonConcreteTargets("src/runsvc/monetdb.tmpl", dataRunSvc, v)

	// 	} else if v == "redisdb" {
	// 		process.ProcessNonConcreteTargets("src/runsvc/redisdb.tmpl", dataRunSvc, v)

	// 	} else if v == "jaeger" {
	// 		process.ProcessNonConcreteTargets("src/runsvc/jaeger.tmpl", dataRunSvc, v)

	// 	}

	// }
}
