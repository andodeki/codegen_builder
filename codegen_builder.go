package main

// https://betterprogramming.pub/how-to-generate-code-using-golang-templates-942cba2e5e0c
import (
	"fmt"

	// "github.com/andodeki/gen_project/builder/codegen_builder/config"
	"github.com/andodeki/gen_project/builder/codegen_builder/src/config"
	"github.com/andodeki/gen_project/builder/codegen_builder/src/data"
	"github.com/andodeki/gen_project/builder/codegen_builder/src/datasources"
	"github.com/andodeki/gen_project/builder/codegen_builder/src/models"
	"github.com/andodeki/gen_project/builder/codegen_builder/src/utils"
)

//  https://go.dev/play/p/zF6y6fWTraq
// const fmtJson = `{"ProductTarget":"%s","ProductSource":"propertylisting","ProductFileType":".sql","ConcreteTargets":["normal","igloo"],"Properties":[{"Name":"windowType","TypeName":"UUID","Tag":""},{"Name":"doorType","TypeName":"TEXT","Tag":""},{"Name":"floor","TypeName":"TIMESTAMP","Tag":""}]}`
// /home/godev/go/src/github.com/andodeki/gen_project/codegen_builder
func main() {

	projectDefaults := data.Data{
		ProductSource: "propertylisting",
		// Workfile:      "../../codegen_builder/", sqlfiles outside this project folder used to generate a project
		Workfile: "./src/sqlfiles/", //sqlfiles within the project folder
		// Workfile: "../../../api.propertylist.com/datasources/migrations/",
		// Workfile:      "../../../api.propertylist.com/datasources/migrations/NEXT/",
		Company:       "andodeki",
		ProjectFolder: "../../tmp", //where the gerated project files will be
	}
	models.ModelsSqlGenerator(projectDefaults)

	dataYAML, kvdbs, dbs, _ := config.YAMLGenerator(projectDefaults) //.GenerateYAML(propts, enums, idxs)

	config.ConfigGenerator(dataYAML, projectDefaults)

	datasources.DataSourcesTargets(kvdbs, dbs, dataYAML, projectDefaults)
	utils.GenerateUtils(projectDefaults)

	fmt.Println("Remember to edit the files that contain the Concrete Targets!")
}
