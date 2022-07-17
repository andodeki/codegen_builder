package models

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"strings"

	"github.com/andodeki/gen_project/builder/codegen_builder/src/data"
	"github.com/andodeki/gen_project/builder/codegen_builder/src/process"
	"github.com/andodeki/gen_project/builder/codegen_builder/src/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ModelsSqlGenerator(projectDefaults data.Data) {
	// propt := data.Property{}
	propts := []data.Property{}
	// enu := data.Enum{}
	enums := []data.Enum{}
	// idx := data.DBIndex{}
	idxs := []data.DBIndex{}
	dataSql := data.Data{
		ProductSource:   projectDefaults.ProductSource,
		Company:         projectDefaults.Company,
		ProductFileType: ".cql",
		ConcreteTargets: []string{},
		FolderPathName:  projectDefaults.ProjectFolder + "/src/datasources/migrations/cqlmigrations/",
		Properties:      propts,
		Enums:           enums,
		DBIndexes:       idxs,
	}
	dataModel := data.Data{

		ProductSource:   projectDefaults.ProductSource,
		Company:         projectDefaults.Company,
		ProductFileType: ".go",
		FolderPathName:  projectDefaults.ProjectFolder + "/src/models/",

		Properties: propts,
		Enums:      enums,
		DBIndexes:  idxs,
	}

	extract := &data.ExtractedData{
		TableName:  "",
		TypeName:   "",
		IndexName:  "",
		Wrkfile:    projectDefaults.Workfile,
		Vals:       []string{"TABLE ", "TYPE ", "INDEX "},
		JsonString: [][]string{},
		WrkFolder:  utils.ReadFolder(projectDefaults.Workfile),
	}
	// wrkFolder := utils.ReadFolder(projectDefaults.Workfile)

	for l, f := range extract.WrkFolder {
		if strings.Contains(f.Name(), ".sql") {
			_, _, j := extractFieldFromSqlFile(extract, l, f, dataSql, dataModel)
			_ = j
		} else if !strings.Contains(f.Name(), ".sql") {
			continue

		}
	}
	// dataModel, dataSql, j := extractFieldFromSqlFileBackUp(projectDefaults.Workfile, dataSql, dataModel)
	// return propts, enums, idxs
}

func extractFieldFromSqlFile(extract *data.ExtractedData, l int, f fs.FileInfo, dataSql data.Data, dataModel data.Data) (data.Data, data.Data, [][]string) {
	content, err := ioutil.ReadFile(extract.Wrkfile + f.Name())
	if err != nil {
		log.Fatal(err)
	}
	extract.JsonString = append(extract.JsonString, strings.Split(string(content), ";")) //= strings.Split(string(content), ";")
	for i, j := range extract.JsonString {
		for w, v := range j {
			extract.JsonString[i][w] = v + ";"
		}
	}

	lines := strings.Split(strings.Join(extract.JsonString[l], "#"), "\n")

	for i, v := range lines {
		if strings.Contains(v, "--") {
			lines[i] = strings.Replace(v, string(v[strings.Index(v, "--")]), "\n --", -1)

		}
	}

	lines = strings.Split(strings.Join(lines, "\n"), "\n")

	for i, line := range lines {
		if strings.Contains(line, "--") {
			lines[i] = ""
		}
	}

	lines = utils.Delete_empty(lines)

	var delNewLines []string
	for _, v := range lines {
		delNewLines = append(delNewLines, strings.TrimSpace(strings.Replace(v, "\t", " ", -1)))
	}
	extract.JsonString[l] = strings.Split(strings.Join(delNewLines, "@"), "#")

	for i, v := range extract.JsonString[l] {
		if strings.Contains(v, "@)") {
			extract.JsonString[l][i] = strings.Replace(v, "@)", ")", -1)
		}
	}
	for i, v := range extract.JsonString[l] {
		if strings.Contains(v, "(@") {
			extract.JsonString[l][i] = strings.Replace(v, "(@", "(", -1)
		}
	}
	for i, v := range extract.JsonString[l] {
		if strings.Contains(v, ",@") {
			extract.JsonString[l][i] = strings.Replace(v, ",@", ", ", -1)
		}
	}
	for i, v := range extract.JsonString[l] {
		if strings.Contains(v, "@") {
			extract.JsonString[l][i] = strings.Replace(v, "@", " ", -1)
		}
	}
	var tableFinal []string
	var typeFields []string
	var linesType []string
	_ = linesType
	var indexFields []string
	tableMap := make(map[string][][]string)
	tableMapNoSubSlice := make(map[string][]string)
	typeMap := make(map[string][]string)
	indexMap := make(map[string][]string)

	for _, jpos := range extract.JsonString[l] {
		for _, pos := range extract.Vals {
			switch pos {
			case "TABLE ":
				if strings.Contains(jpos, "TABLE ") {
					var tableFields2 []string
					var tableLineFinal2 [][]string
					var linestbl []string
					ls := utils.Delete_empty(strings.Split(jpos, " "))
					for i := range ls {
						if ls[i] == strings.TrimSpace(pos) {
							extract.TableName = utils.Before(strings.TrimSpace(ls[i+1]), "(") //strings.Replace(strings.TrimSpace(ls[i+1]), "(", "", -1)
						}
						if strings.Contains(ls[i], extract.TableName+"(") { //ls[i] == strings.TrimSpace("ENUM(") { //ls[i+len(ls)]
							_ = tableFields2
							ll := (utils.After(strings.TrimSpace(strings.Join(ls, " ")), extract.TableName+"("))
							if strings.Contains(ll, "PRIMARY KEY(") {
								ll = utils.Before(utils.After(strings.TrimSpace(strings.Join(ls, " ")), extract.TableName+"("), "PRIMARY KEY(")
							}
							if strings.Contains(ll, "PRIMARY KEY (") {
								ll = utils.Before(utils.After(strings.TrimSpace(strings.Join(ls, " ")), extract.TableName+"("), "PRIMARY KEY (")
							}

							tableFields2 = append(tableFields2, ll)
							linestbl = strings.Split(ll, ",")
							for i, j := range linestbl {
								linestbl[i] = j + ","
							}
							linestbl = tableFields2
							tableLineFinal2 = append(tableLineFinal2, linestbl)
							tableMap[extract.TableName] = tableLineFinal2
							tableMapNoSubSlice[extract.TableName] = tableFields2
						}
					}
				}
			case "TYPE ":
				if strings.Contains(jpos, "TYPE ") {
					ls := utils.Delete_empty(strings.Split(jpos, " "))
					for i := range ls {
						if ls[i] == strings.TrimSpace(pos) {
							if strings.Contains(strings.TrimSpace(ls[i+1]), "(") {
								extract.TypeName = utils.Before(strings.TrimSpace(ls[i+1]), "(") //strings.Replace(strings.TrimSpace(ls[i+1]), "(", "", -1)
							}
							extract.TypeName = strings.TrimSpace(ls[i+1]) //strings.Replace(strings.TrimSpace(ls[i+1]), "(", "", -1)
						}
						var linestye []string
						if strings.Contains(ls[i], "ENUM") { //ls[i] == strings.TrimSpace("ENUM(") { //ls[i+len(ls)]
							for j := range ls[i:] {
								if !strings.Contains(ls[i], ");") {
									typeFields = append(typeFields, strings.TrimSpace(ls[i+j]))
								} else if strings.Contains(ls[i], "ENUM") {
									typeFields = append(typeFields, strings.TrimSpace(ls[i]))
								}
							}
							linestye = strings.Split((utils.After(strings.Join(typeFields, " "), "ENUM")), ",")
							for i, j := range linestye {
								linestye[i] = j + ","
							}
							typeMap[extract.TypeName] = linestye
						}
					}
				}
			case "INDEX ":
				if strings.Contains(jpos, "INDEX ") {
					ls := utils.Delete_empty(strings.Split(jpos, " "))
					for i := range ls {
						if ls[i] == strings.TrimSpace(pos) {
							if strings.Contains(strings.TrimSpace(ls[i+1]), "(") {
								extract.IndexName = utils.Before(strings.TrimSpace(ls[i+1]), "(") //strings.Replace(strings.TrimSpace(ls[i+1]), "(", "", -1)
							}
							extract.IndexName = strings.TrimSpace(ls[i+1])     //strings.Replace(strings.TrimSpace(ls[i+1]), "(", "", -1)
							idxname := data.IndexName{Name: extract.IndexName} //data.DBIndex{Name: extract.IndexName}
							dataSql.ProductTargets.IndexNames = append(dataSql.ProductTargets.IndexNames, idxname)
						}
						if ls[i] == strings.TrimSpace("ON") {
							for j, _ := range ls[i:] {
								if !strings.Contains(ls[i], ");") {
									indexFields = append(indexFields, strings.TrimSpace(ls[i+j]))
								} else if strings.Contains(ls[i], "ON") {
									indexFields = append(indexFields, strings.TrimSpace(ls[i]))
								}
							}
							linesidx := strings.Split((utils.After(strings.Join(indexFields, " "), "ON")), ",")
							for i, j := range linesidx {
								linesidx[i] = j + ","
							}
							indexMap[extract.IndexName] = linesidx
						}
					}
				}
			}
		}
	}

	tableMapWithSubSlice := make(map[string][][]string)
	for key, value := range tableMapNoSubSlice {
		var sl []string
		var tableLineFinal [][]string
		for _, v := range value {
			for _, w := range strings.Split(v, ",") {
				sl = utils.Delete_empty(strings.Split(w, " "))
				if len(sl) > 1 {
					tableLineFinal = append(tableLineFinal, sl[:2])
				}
			}
		}
		tableMapWithSubSlice[key] = tableLineFinal
	}
	_ = tableFinal
	pt := new(data.ProductTarget)
	pt.Structs = make(map[data.TableName][]data.Property)
	dmpt := new(data.ProductTarget)
	dmpt.Structs = make(map[data.TableName][]data.Property)
	var nameSql string
	for key, values := range tableMapWithSubSlice {

		tblname := data.TableName{Name: extract.TableName}
		// dataSql.FolderPathName = "./tmp/datasources/migrations/cqlmigrations/"
		dataSql.ProductFileType = ".cql"
		if strings.Contains(f.Name(), ".up.sql") {
			dataSql.ProductTargets.OutputFilename = []string{utils.Before(f.Name()[7:], "_table.up.sql")}
		}
		if strings.Contains(f.Name(), ".down.sql") {
			dataSql.ProductTargets.OutputFilename = []string{utils.Before(f.Name()[7:], ".down.sql")}
		}
		nameSql = dataSql.ProductTargets.OutputFilename[0]

		var properties []data.Property

		if dataSql.ProductFileType == ".cql" {
			for _, value := range values {
				prop := data.Property{
					Name:     value[0],
					TypeName: utils.ToOutPutDataType(utils.RegxReplaceAllString(value[1], "@"), dataSql),
					Tag:      "",
				}
				properties = append(properties, prop)
			}

			tblname = data.TableName{key, dataSql.ProductSource}
			pt.Structs[tblname] = properties
			dataSql.AddProductTargetProperty(*pt)
		}
	}

	// fmt.Printf("pt: %q\n", pt.Structs)

	jsonStr, _ := json.Marshal(tableMapWithSubSlice)
	var results []map[string]interface{}

	json.Unmarshal([]byte("["+string(jsonStr)+"]"), &results)

	for _, value := range typeMap {
		for i, v := range value {
			ls := utils.RegxReplaceAllString(strings.Join(utils.Delete_empty(strings.Split(strings.Replace(v, "(", "", -1), " ")), ""), "")
			value[i] = ls
		}
	}

	pty := new(data.ProductTarget)
	pty.Types = make(map[string][]data.Enum, 0)
	for key, value := range typeMap {
		if dataSql.ProductFileType == ".cql" {
			var enums []data.Enum
			for _, vals := range value {
				if len(vals) > 0 {
					enum := data.Enum{Name: vals}
					enums = append(enums, enum)
				}
			}
			pty.Types[key] = enums
			dataSql.AddProductTargetTypes(*pty)
		}
	}
	// fmt.Printf("typeMap: %v\n", typeMap)
	for _, value := range indexMap {
		for i, v := range value {
			li := utils.RegxReplaceAllString(v, "_(")
			value[i] = li
		}
	}

	pidx := new(data.ProductTarget)
	pidx.Index = make(map[data.IndexName][]data.DBIndex, 0)

	for key, value := range indexMap {
		var indexes []data.DBIndex
		var idxname data.IndexName //data.TableName
		for _, vals := range value {
			if len(vals) > 0 {
				if strings.Contains(vals, extract.TableName+"(") {
					index := data.DBIndex{Name: utils.After(vals, extract.TableName+"("), TableName: extract.TableName}
					indexes = append(indexes, index)
				} else {
					index := data.DBIndex{Name: vals, TableName: extract.TableName}
					indexes = append(indexes, index)
				}
			}
		}
		idxname = data.IndexName{key, extract.TableName}
		pidx.Index[idxname] = indexes
	}

	dataSql.AddProductTargetIndex(*pidx)

	var nameModel string
	for key, values := range tableMapWithSubSlice {
		tblname := data.TableName{Name: extract.TableName}
		dataModel.ProductTargets.TableNames = append(dataModel.ProductTargets.TableNames, tblname)
		dataModel.ProductTargets.StructNames = append(dataModel.ProductTargets.StructNames, tblname)
		if strings.Contains(f.Name(), ".up.sql") {
			dataModel.ProductTargets.OutputFilename = []string{utils.Before(f.Name()[7:], "_table.up.sql")}
		}
		if strings.Contains(f.Name(), ".down.sql") {
			dataModel.ProductTargets.OutputFilename = []string{utils.Before(f.Name()[7:], ".down.sql")}
		}
		nameModel = dataModel.ProductTargets.OutputFilename[0]
		dataModel.ProductFileType = ".go"
		// dataModel.FolderPathName = "./tmp/models/"
		var properties []data.Property
		if dataModel.ProductFileType == ".go" {
			// dmpt := new(data.ProductTarget)
			// dmpt.Structs = make(map[data.TableName][]data.Property)
			for i, value := range values {
				// fmt.Printf("value: %v\n", value)
				prop := data.Property{
					Name:     strings.Replace(cases.Title(language.Und, cases.NoLower).String(strings.Replace(value[0][2:], "_", " ", -1)), " ", "", -1), //value[0],
					TypeName: utils.ToOutPutDataType(utils.RegxReplaceAllString(value[1], "@_"), dataModel),
					// Tag:      "",
				}
				// fmt.Printf("prop.TypeName: %v\n", prop.TypeName)
				// fmt.Printf("i: %v\n", i)
				if i == 0 {
					prop.Tag = fmt.Sprintf(`json:"%s,omitempty" db:"%s"`, prop.Name, value[0])
				} else {
					prop.Tag = fmt.Sprintf(`json:"%s" db:"%s"`, strings.ToLower(prop.Name), value[0])
				}
				properties = append(properties, prop)
			}

			tblname = data.TableName{key, dataModel.ProductSource}
			dmpt.Structs[tblname] = properties
			dataModel.AddProductTargetProperty(*dmpt)
		}
	}

	dmpty := new(data.ProductTarget)
	dmpty.Types = make(map[string][]data.Enum, 0)
	for key, value := range typeMap {
		if dataModel.ProductFileType == ".go" {
			var enums []data.Enum

			for _, vals := range value {

				if len(vals) > 0 {
					str := strings.Replace(
						cases.Title(
							language.Und, cases.NoLower,
						).String(
							strings.Replace(key, "_", " ", -1),
						), " ", "", -1,
					)
					enum := data.Enum{Name: vals, CustomType: str}
					enums = append(enums, enum)
				}
			}
			dmpty.Types[key] = enums
			dataModel.AddProductTargetTypes(*dmpty)
		}
	}
	newFunction(dataSql, dataModel, nameModel, nameSql)

	// } else if !strings.Contains(f.Name(), ".sql") {
	// 	// continue

	// }

	// // }

	return dataModel, dataSql, extract.JsonString
}

func newFunction(dataSql data.Data, dataModel data.Data, nameModel string, nameSql string) {

	for _, v := range dataSql.ProductTargets.OutputFilename {
		fmt.Printf("v: %v\n", v)
		if v == nameSql {
			process.ProcessNonConcreteTargets("src/models/DBProductTarget.tmpl", dataSql, v)
			dataSql.ClearProductTargetProperty()

			dataSql.ClearProductTargetTypes()
			dataSql.ClearProductTargetIndex()
		}
	}

	for _, v := range dataModel.ProductTargets.OutputFilename {
		if v == nameModel {
			process.ProcessNonConcreteTargets("src/models/Models.tmpl", dataModel, v)
			dataModel.ClearProductTargetProperty()

			dataModel.ClearProductTargetIndex()
		}
	}

}

func extractFieldFromSqlFileBackUp(wrkfile string, dataSql data.Data, dataModel data.Data) (data.Data, data.Data, [][]string) {
	wrkFolder := utils.ReadFolder(wrkfile)

	var tableName, typeName, indexName string

	vals := []string{"TABLE ", "TYPE ", "INDEX "}

	var jsonString [][]string
	for l, f := range wrkFolder {
		if strings.Contains(f.Name(), ".sql") {
			content, err := ioutil.ReadFile(wrkfile + f.Name())
			if err != nil {
				log.Fatal(err)
			}
			jsonString = append(jsonString, strings.Split(string(content), ";")) //= strings.Split(string(content), ";")
			for i, j := range jsonString {
				for w, v := range j {
					jsonString[i][w] = v + ";"
				}
			}

			lines := strings.Split(strings.Join(jsonString[l], "#"), "\n")

			for i, v := range lines {
				if strings.Contains(v, "--") {
					lines[i] = strings.Replace(v, string(v[strings.Index(v, "--")]), "\n --", -1)

				}
			}

			lines = strings.Split(strings.Join(lines, "\n"), "\n")

			for i, line := range lines {
				if strings.Contains(line, "--") {
					lines[i] = ""
				}
			}

			lines = utils.Delete_empty(lines)

			var delNewLines []string
			for _, v := range lines {
				delNewLines = append(delNewLines, strings.TrimSpace(strings.Replace(v, "\t", " ", -1)))
			}
			jsonString[l] = strings.Split(strings.Join(delNewLines, "@"), "#")

			for i, v := range jsonString[l] {
				if strings.Contains(v, "@)") {
					jsonString[l][i] = strings.Replace(v, "@)", ")", -1)
				}
			}
			for i, v := range jsonString[l] {
				if strings.Contains(v, "(@") {
					jsonString[l][i] = strings.Replace(v, "(@", "(", -1)
				}
			}
			for i, v := range jsonString[l] {
				if strings.Contains(v, ",@") {
					jsonString[l][i] = strings.Replace(v, ",@", ", ", -1)
				}
			}
			for i, v := range jsonString[l] {
				if strings.Contains(v, "@") {
					jsonString[l][i] = strings.Replace(v, "@", " ", -1)
				}
			}
			var tableFinal []string
			var typeFields []string
			var linesType []string
			_ = linesType
			var indexFields []string
			tableMap := make(map[string][][]string)
			tableMapNoSubSlice := make(map[string][]string)
			typeMap := make(map[string][]string)
			indexMap := make(map[string][]string)

			for _, jpos := range jsonString[l] {
				for _, pos := range vals {
					switch pos {
					case "TABLE ":
						if strings.Contains(jpos, "TABLE ") {
							var tableFields2 []string
							var tableLineFinal2 [][]string
							var linestbl []string
							ls := utils.Delete_empty(strings.Split(jpos, " "))
							for i := range ls {
								if ls[i] == strings.TrimSpace(pos) {
									tableName = utils.Before(strings.TrimSpace(ls[i+1]), "(") //strings.Replace(strings.TrimSpace(ls[i+1]), "(", "", -1)
								}
								if strings.Contains(ls[i], tableName+"(") { //ls[i] == strings.TrimSpace("ENUM(") { //ls[i+len(ls)]
									_ = tableFields2
									ll := (utils.After(strings.TrimSpace(strings.Join(ls, " ")), tableName+"("))
									if strings.Contains(ll, "PRIMARY KEY(") {
										ll = utils.Before(utils.After(strings.TrimSpace(strings.Join(ls, " ")), tableName+"("), "PRIMARY KEY(")
									}
									if strings.Contains(ll, "PRIMARY KEY (") {
										ll = utils.Before(utils.After(strings.TrimSpace(strings.Join(ls, " ")), tableName+"("), "PRIMARY KEY (")
									}

									tableFields2 = append(tableFields2, ll)
									linestbl = strings.Split(ll, ",")
									for i, j := range linestbl {
										linestbl[i] = j + ","
									}
									linestbl = tableFields2
									tableLineFinal2 = append(tableLineFinal2, linestbl)
									tableMap[tableName] = tableLineFinal2
									tableMapNoSubSlice[tableName] = tableFields2
								}
							}
						}
					case "TYPE ":
						if strings.Contains(jpos, "TYPE ") {
							ls := utils.Delete_empty(strings.Split(jpos, " "))
							for i := range ls {
								if ls[i] == strings.TrimSpace(pos) {
									if strings.Contains(strings.TrimSpace(ls[i+1]), "(") {
										typeName = utils.Before(strings.TrimSpace(ls[i+1]), "(") //strings.Replace(strings.TrimSpace(ls[i+1]), "(", "", -1)
									}
									typeName = strings.TrimSpace(ls[i+1]) //strings.Replace(strings.TrimSpace(ls[i+1]), "(", "", -1)
								}
								var linestye []string
								if strings.Contains(ls[i], "ENUM") { //ls[i] == strings.TrimSpace("ENUM(") { //ls[i+len(ls)]
									for j := range ls[i:] {
										if !strings.Contains(ls[i], ");") {
											typeFields = append(typeFields, strings.TrimSpace(ls[i+j]))
										} else if strings.Contains(ls[i], "ENUM") {
											typeFields = append(typeFields, strings.TrimSpace(ls[i]))
										}
									}
									linestye = strings.Split((utils.After(strings.Join(typeFields, " "), "ENUM")), ",")
									for i, j := range linestye {
										linestye[i] = j + ","
									}
									typeMap[typeName] = linestye
								}
							}
						}
					case "INDEX ":
						if strings.Contains(jpos, "INDEX ") {
							ls := utils.Delete_empty(strings.Split(jpos, " "))
							for i := range ls {
								if ls[i] == strings.TrimSpace(pos) {
									if strings.Contains(strings.TrimSpace(ls[i+1]), "(") {
										indexName = utils.Before(strings.TrimSpace(ls[i+1]), "(") //strings.Replace(strings.TrimSpace(ls[i+1]), "(", "", -1)
									}
									indexName = strings.TrimSpace(ls[i+1])     //strings.Replace(strings.TrimSpace(ls[i+1]), "(", "", -1)
									idxname := data.IndexName{Name: indexName} //data.DBIndex{Name: indexName}
									dataSql.ProductTargets.IndexNames = append(dataSql.ProductTargets.IndexNames, idxname)
								}
								if ls[i] == strings.TrimSpace("ON") {
									for j, _ := range ls[i:] {
										if !strings.Contains(ls[i], ");") {
											indexFields = append(indexFields, strings.TrimSpace(ls[i+j]))
										} else if strings.Contains(ls[i], "ON") {
											indexFields = append(indexFields, strings.TrimSpace(ls[i]))
										}
									}
									linesidx := strings.Split((utils.After(strings.Join(indexFields, " "), "ON")), ",")
									for i, j := range linesidx {
										linesidx[i] = j + ","
									}
									indexMap[indexName] = linesidx
								}
							}
						}
					}
				}
			}

			tableMapWithSubSlice := make(map[string][][]string)
			for key, value := range tableMapNoSubSlice {
				var sl []string
				var tableLineFinal [][]string
				for _, v := range value {
					for _, w := range strings.Split(v, ",") {
						sl = utils.Delete_empty(strings.Split(w, " "))
						if len(sl) > 1 {
							tableLineFinal = append(tableLineFinal, sl[:2])
						}
					}
				}
				tableMapWithSubSlice[key] = tableLineFinal
			}
			_ = tableFinal
			pt := new(data.ProductTarget)
			pt.Structs = make(map[data.TableName][]data.Property)
			dmpt := new(data.ProductTarget)
			dmpt.Structs = make(map[data.TableName][]data.Property)
			var nameSql string
			for key, values := range tableMapWithSubSlice {

				tblname := data.TableName{Name: tableName}
				// dataSql.FolderPathName = "./tmp/datasources/migrations/cqlmigrations/"
				dataSql.ProductFileType = ".cql"
				if strings.Contains(f.Name(), ".up.sql") {
					dataSql.ProductTargets.OutputFilename = []string{utils.Before(f.Name()[7:], "_table.up.sql")}
				}
				if strings.Contains(f.Name(), ".down.sql") {
					dataSql.ProductTargets.OutputFilename = []string{utils.Before(f.Name()[7:], ".down.sql")}
				}
				nameSql = dataSql.ProductTargets.OutputFilename[0]

				var properties []data.Property

				if dataSql.ProductFileType == ".cql" {
					for _, value := range values {
						prop := data.Property{
							Name:     value[0],
							TypeName: utils.ToOutPutDataType(utils.RegxReplaceAllString(value[1], "@"), dataSql),
							Tag:      "",
						}
						properties = append(properties, prop)
					}

					tblname = data.TableName{key, dataSql.ProductSource}
					pt.Structs[tblname] = properties
					dataSql.AddProductTargetProperty(*pt)
				}
			}

			// fmt.Printf("pt: %q\n", pt.Structs)

			jsonStr, _ := json.Marshal(tableMapWithSubSlice)
			var results []map[string]interface{}

			json.Unmarshal([]byte("["+string(jsonStr)+"]"), &results)

			for _, value := range typeMap {
				for i, v := range value {
					ls := utils.RegxReplaceAllString(strings.Join(utils.Delete_empty(strings.Split(strings.Replace(v, "(", "", -1), " ")), ""), "")
					value[i] = ls
				}
			}

			pty := new(data.ProductTarget)
			pty.Types = make(map[string][]data.Enum, 0)
			for key, value := range typeMap {
				if dataSql.ProductFileType == ".cql" {
					var enums []data.Enum
					for _, vals := range value {
						if len(vals) > 0 {
							enum := data.Enum{Name: vals}
							enums = append(enums, enum)
						}
					}
					pty.Types[key] = enums
					dataSql.AddProductTargetTypes(*pty)
				}
			}
			// fmt.Printf("typeMap: %v\n", typeMap)
			for _, value := range indexMap {
				for i, v := range value {
					li := utils.RegxReplaceAllString(v, "_(")
					value[i] = li
				}
			}

			pidx := new(data.ProductTarget)
			pidx.Index = make(map[data.IndexName][]data.DBIndex, 0)

			for key, value := range indexMap {
				var indexes []data.DBIndex
				var idxname data.IndexName //data.TableName
				for _, vals := range value {
					if len(vals) > 0 {
						if strings.Contains(vals, tableName+"(") {
							index := data.DBIndex{Name: utils.After(vals, tableName+"("), TableName: tableName}
							indexes = append(indexes, index)
						} else {
							index := data.DBIndex{Name: vals, TableName: tableName}
							indexes = append(indexes, index)
						}
					}
				}
				idxname = data.IndexName{key, tableName}
				pidx.Index[idxname] = indexes
			}

			dataSql.AddProductTargetIndex(*pidx)

			var nameModel string
			for key, values := range tableMapWithSubSlice {
				tblname := data.TableName{Name: tableName}
				dataModel.ProductTargets.TableNames = append(dataModel.ProductTargets.TableNames, tblname)
				dataModel.ProductTargets.StructNames = append(dataModel.ProductTargets.StructNames, tblname)
				if strings.Contains(f.Name(), ".up.sql") {
					dataModel.ProductTargets.OutputFilename = []string{utils.Before(f.Name()[7:], "_table.up.sql")}
				}
				if strings.Contains(f.Name(), ".down.sql") {
					dataModel.ProductTargets.OutputFilename = []string{utils.Before(f.Name()[7:], ".down.sql")}
				}
				nameModel = dataModel.ProductTargets.OutputFilename[0]
				dataModel.ProductFileType = ".go"
				// dataModel.FolderPathName = "./tmp/models/"
				var properties []data.Property
				if dataModel.ProductFileType == ".go" {
					// dmpt := new(data.ProductTarget)
					// dmpt.Structs = make(map[data.TableName][]data.Property)
					for i, value := range values {
						// fmt.Printf("value: %v\n", value)
						prop := data.Property{
							Name:     strings.Replace(cases.Title(language.Und, cases.NoLower).String(strings.Replace(value[0][2:], "_", " ", -1)), " ", "", -1), //value[0],
							TypeName: utils.ToOutPutDataType(utils.RegxReplaceAllString(value[1], "@_"), dataModel),
							// Tag:      "",
						}
						// fmt.Printf("prop.TypeName: %v\n", prop.TypeName)
						// fmt.Printf("i: %v\n", i)
						if i == 0 {
							prop.Tag = fmt.Sprintf(`json:"%s,omitempty" db:"%s"`, prop.Name, value[0])
						} else {
							prop.Tag = fmt.Sprintf(`json:"%s" db:"%s"`, strings.ToLower(prop.Name), value[0])
						}
						properties = append(properties, prop)
					}

					tblname = data.TableName{key, dataModel.ProductSource}
					dmpt.Structs[tblname] = properties
					dataModel.AddProductTargetProperty(*dmpt)
				}
			}

			dmpty := new(data.ProductTarget)
			dmpty.Types = make(map[string][]data.Enum, 0)
			for key, value := range typeMap {
				if dataModel.ProductFileType == ".go" {
					var enums []data.Enum

					for _, vals := range value {

						if len(vals) > 0 {
							str := strings.Replace(
								cases.Title(
									language.Und, cases.NoLower,
								).String(
									strings.Replace(key, "_", " ", -1),
								), " ", "", -1,
							)
							enum := data.Enum{Name: vals, CustomType: str}
							enums = append(enums, enum)
						}
					}
					dmpty.Types[key] = enums
					dataModel.AddProductTargetTypes(*dmpty)
				}
			}
			newFunction(dataSql, dataModel, nameModel, nameSql)

		} else if !strings.Contains(f.Name(), ".sql") {
			continue
		}

	}

	return dataModel, dataSql, jsonString
}
