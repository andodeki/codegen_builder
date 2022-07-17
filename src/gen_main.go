package src

import (
	"strings"

	"github.com/andodeki/gen_project/builder/codegen_builder/src/data"
	"github.com/andodeki/gen_project/builder/codegen_builder/src/process"
)

func InMain(projectDefaults data.Data) {
	cts := make([]string, 0)
	addFiles := []string{"main", "startup"} //adding extra files not in configuration yaml file
	for _, v := range addFiles {
		cts = append(cts, strings.ToLower(strings.Replace(v, "Clients", "", -1)))
	}
	// fmt.Printf("cts: %v\n", cts)
	dataRunSvc := data.Data{
		ProductTargets: data.ProductTarget{
			OutputFilename: cts,
			DSource: data.Datasource{
				// DBSMap:       kvdbs,
				// DBSMapFields: dsrcs.DSource.DBSMapFields,
				DBS: addFiles,
				// IMethods:     iterfacemethods,
			},
			// ProductFileTypeFrom: dataYAML.ProductFileType,
		},
		ProductSource:   projectDefaults.ProductSource,
		Company:         projectDefaults.Company,
		ProductFileType: ".go",
		FolderPathName:  projectDefaults.ProjectFolder + "/src/",
		ConcreteTargets: cts,
	}
	process.ProcessConcreteTargets("src/main.tmpl", dataRunSvc)
	process.ProcessConcreteTargets("src/startup.tmpl", dataRunSvc)
	//[httpserver scylladb pasetotoken monetdb jaeger configuration postgresdb redisdb email queue]
	// for _, v := range dataRunSvc.ProductTargets.OutputFilename {
	// 	if v == "main" {
	// 		process.ProcessNonConcreteTargets("src/main.tmpl", dataRunSvc, v)

	// 	} else if v == "startup" {
	// 		process.ProcessNonConcreteTargets("src/startup.tmpl", dataRunSvc, v)

	// 	}
	// }
}
