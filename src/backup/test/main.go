package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	// "golang.org/x/text/cases"
	// "golang.org/x/text/language"

	"github.com/Masterminds/sprig"
)

type Data struct {
	ProductTargets  ProductTarget
	ProductSource   string
	ProductFileType string
	ConcreteTargets []string
	Properties      []Property
	Enums           []Enum
	DBIndexes       []DBIndex
}

type ProductTarget struct {
	StructName string
	TableName  string
	TypeName   string
	IndexName  string
}

type Property struct {
	Name     string
	TypeName string
	Tag      string
}

type Enum struct {
	Name string
}

type DBIndex struct {
	Name string
}

func (prop *Data) AddProperty(p Property) []Property {
	prop.Properties = append(prop.Properties, p)
	return prop.Properties
}
func (e *Data) AddEnum(p Enum) []Enum {
	e.Enums = append(e.Enums, p)
	return e.Enums
}

func (i *Data) AddDBIndex(p DBIndex) []DBIndex {
	i.DBIndexes = append(i.DBIndexes, p)
	return i.DBIndexes
}
func main() {
	propt := Property{}
	propts := []Property{}
	enu := Enum{}
	enums := []Enum{}
	idx := DBIndex{}
	idxs := []DBIndex{}
	dataSql := Data{
		// ProductTargets:  ProductTarget{tableName, tableName, typeName, indexName},
		ProductSource:   "propertylisting",
		ProductFileType: ".cql",
		ConcreteTargets: []string{},
		Properties:      propts,
		Enums:           enums,
		DBIndexes:       idxs,
	}
	dataModel := Data{
		// ProductTargets:  ProductTarget{tableName[:len(tableName)-1], tableName, typeName, indexName},
		ProductSource:   "propertylisting",
		ProductFileType: ".go",
		// ConcreteTargets: []string{tableName[:len(tableName)-1]},
		Properties: propts,
		Enums:      enums,
		DBIndexes:  idxs,
	}

	var wrkfile = "../../codegen_builder/"
	dataModel, dataSql, j := extractFieldFromSqlFile(wrkfile, propt, enu, idx, dataSql, dataModel)
	for _, v := range j {
		fmt.Printf("v: %v\n", len(v))

		// fmt.Printf("v////=====================================================\n: %v\n", v)
	}
	// fmt.Printf("============================================================j: %v\n", j)
}

func extractFieldFromSqlFile(wrkfile string, propt Property, enu Enum, idx DBIndex, dataSql Data, dataModel Data) (Data, Data, [][]string) {
	wrkFolder := readFolder(wrkfile)

	var tableName string
	// var tableFields []string
	// var tableFieldsSubSlice [][]string
	var SubSliceTables [][]string
	// var linesTable [][]string
	vals := []string{"TABLE ", "TYPE ", "INDEX "}

	var jsonString [][]string
	for l, f := range wrkFolder {
		if strings.Contains(f.Name(), ".sql") {
			content, err := ioutil.ReadFile(wrkfile + f.Name())
			if err != nil {
				log.Fatal(err)
			}
			jsonString = append(jsonString, strings.Split(string(content), ";")) //= strings.Split(string(content), ";")
			// vals := []string{"TABLE ", "TYPE ", "INDEX "}
			for i, j := range jsonString {
				for w, v := range j {
					jsonString[i][w] = v + ";"
				}
				// fmt.Printf("jsonString[i][w]: %v\n", len(jsonString[i]))
			}
			lines := strings.Split(strings.Join(jsonString[l], "#"), "\n")
			for i, line := range lines {
				if strings.Contains(line, "--") {
					lines[i] = ""
				}
			}

			// fmt.Printf("lines: %v\n", len(lines))
			lines = delete_empty(lines)

			var delNewLines []string
			for _, v := range lines {
				delNewLines = append(delNewLines, strings.TrimSpace(strings.Replace(v, "\t", " ", -1)))
				// fmt.Printf("v: %q\n", strings.Replace(v, "\t", " ", -1))
			}
			jsonString[l] = strings.Split(strings.Join(delNewLines, "@"), "#")

			// fmt.Printf("jsonString: %v\n", jsonString[l])
			// fmt.Printf("lines1: %v\n", len(lines))

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
			// fmt.Printf("jsonString: %v\n", (jsonString[l]))
			for _, jpos := range jsonString[l] {
				for _, pos := range vals {
					switch pos {
					case "TABLE ":
						if strings.Contains(jpos, "TABLE ") {
							// fmt.Printf("jpos{%v}: %v\n", len(jpos), i)

							ls := delete_empty(strings.Split(jpos, " "))
							// fmt.Printf("ls: %v\n", ls)
							SubSliceTables = append(SubSliceTables, ls)
							// fmt.Printf("SubSlice{%v}: %v\n", (SubSliceTables), i)
							for i := range ls {
								if ls[i] == strings.TrimSpace(pos) {
									tableName = before(strings.TrimSpace(ls[i+1]), "(") //strings.Replace(strings.TrimSpace(ls[i+1]), "(", "", -1)
									// fmt.Printf("tableName: %v\n", tableName)
									if dataSql.ProductFileType == ".cql" {
										dataSql.ProductTargets.TableName = tableName
									} else if dataModel.ProductFileType == ".go" {
										// ConcreteTargets: []string{tableName[:len(tableName)-1]},

										dataModel.ProductTargets.StructName = tableName
										dataModel.ConcreteTargets = append(dataModel.ConcreteTargets, tableName)
									}
								} else if strings.Contains(ls[i], tableName+"(") { //ls[i] == strings.TrimSpace("ENUM(") { //ls[i+len(ls)]
									ll := (after(strings.TrimSpace(strings.Join(ls, " ")), tableName+"("))
									if strings.Contains(ll, "PRIMARY KEY(") {
										ll = before(after(strings.TrimSpace(strings.Join(ls, " ")), tableName+"("), "PRIMARY KEY(")
									} else if strings.Contains(ll, "PRIMARY KEY (") {
										ll = before(after(strings.TrimSpace(strings.Join(ls, " ")), tableName+"("), "PRIMARY KEY (")
									}
									// fmt.Printf("ll: %v\n", ll)
									linestbl := strings.Split(ll, ",")
									for i, j := range linestbl {
										linestbl[i] = j + ","
									}
									// fmt.Printf("linestbl: %v\n", (linestbl))
									for _, v := range linestbl {
										if len(v) > 6 {
											lines := delete_empty(strings.Split(strings.Replace(v, "(", "", -1), " "))
											final := strings.Split(strings.Join(lines[:2], " "), " ")
											fmt.Printf("final: %v\n", final)
											if dataSql.ProductFileType == ".cql" {
												propt.Name = final[0]
												propt.TypeName = toOutPutDataType(final[1], dataSql)
												propt.Tag = ""
												dataSql.AddProperty(propt)
											} else if dataModel.ProductFileType == ".go" {
												propt.Name = strings.Replace(cases.Title(language.Und, cases.NoLower).String(strings.Replace(final[0][2:], "_", " ", -1)), " ", "", -1)
												propt.TypeName = toOutPutDataType(final[1], dataModel)
												if i == 0 {
													propt.Tag = fmt.Sprintf(`json:"%s,omitempty" db:"%s"`, propt.Name, final[0])

												} else {
													propt.Tag = fmt.Sprintf(`json:"%s" db:"%s"`, strings.ToLower(propt.Name), final[0])

												}

												dataModel.AddProperty(propt)
											}
										}

									}
								}
							}
						}
					}
				}
			}
		} else if !strings.Contains(f.Name(), ".sql") {
			continue
		}

	}

	return dataModel, dataSql, jsonString
}

func readFolder(filename string) []fs.FileInfo {
	files, err := ioutil.ReadDir(filename)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

// https://www.dotnetperls.com/between-before-after-go
func between(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

func after(value string, a string) string {
	// Get substring after a string.
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:len(value)]
}

func before(value string, a string) string {
	// Get substring before a string.
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}

func toOutPutDataType(ty string, op Data) string {
	if op.ProductFileType == ".cql" {
		switch ty {
		case "UUID":
			return "UUID"
		case "TEXT":
			return "TEXT"
		case "TIMESTAMPTZ":
			return "TIMESTAMP"
		default:
			return ty
		}
		// if ty == "UUID"{}
	} else if op.ProductFileType == ".go" {
		switch ty {
		case "UUID":
			return "*string"
		case "TEXT":
			return "*string"
		case "TIMESTAMPTZ":
			return "*time.Time"
		case "bytea":
			return "*[]byte "
		default:
			return ty
		}
		// if ty == "UUID"{}
	} else {
		return ty
	}
	return ""
}

func processTemplate(fileName string, outputFile string, data Data) {
	var helpers template.FuncMap = map[string]interface{}{
		"isLast": func(index int, len int) bool {
			return index+1 == len
		},
		"stringSlice": func(s string, i, j int) string {
			return s[i:j]
		},
		"inc": func(i int) int {
			return i + 1
		},
		"dec": func(i int) int {
			return i - 1
		},
	}

	// t.Funcs(template.FuncMap{
	// 	"stringSlice": func(s string, i, j int) string {
	// 		return s[i:j]
	// 	}
	// })
	tmpl := template.Must(template.New("").Funcs(sprig.FuncMap()).Funcs(helpers).ParseFiles(fileName))
	var processed bytes.Buffer
	err := tmpl.ExecuteTemplate(&processed, fileName, data)
	if err != nil {
		log.Fatalf("Unable to parse data into template: %v\n", err)
	}

	if data.ProductFileType == ".cql" {
		// sprocessed := processed.String()
		formatted := processed.Bytes()
		outputPath := "./tmp/" + outputFile
		fmt.Println("Writing file: ", outputPath)
		f, _ := os.Create(outputPath)
		w := bufio.NewWriter(f)
		newlineString := strings.Replace(string(formatted), "<br>", "\n", -1)
		tabString := strings.Replace(newlineString, "<brt>", "\t", -1)
		w.WriteString(tabString)
		w.Flush()
	} else if data.ProductFileType == ".go" {
		fmt.Printf("fileName: %v\n", fileName)
		// formatted, err := format.Source(processed.Bytes())
		// if err != nil {
		// 	log.Fatalf("Could not format processed template: %v\n", err)
		// }
		formatted := processed.Bytes()

		outputPath := "./tmp/" + outputFile
		fmt.Println("Writing file: ", outputPath)
		f, _ := os.Create(outputPath)
		w := bufio.NewWriter(f)
		newlineString := strings.Replace(string(formatted), "<br>", "\n", -1)
		tabString := strings.Replace(newlineString, "<brt>", "\t", -1)
		rmQuotes := strings.Replace(tabString, "&#34;", "\"", -1)
		w.WriteString(rmQuotes)
		w.Flush()
	}

}

func processConcreteTargets(fileName string, data Data) {
	if data.ProductFileType == ".cql" {
		// for _, value := range data.ProductTarget {
		// value := data.ProductTarget
		newData := data
		newData.ProductTargets = data.ProductTargets
		processTemplate(fileName, data.ProductTargets.TableName+".cql", newData)
		// }
	} else if data.ProductFileType == ".go" {
		// fmt.Printf("fileName: %v\n", fileName)
		for _, value := range data.ConcreteTargets {
			newData := data
			newData.ConcreteTargets = []string{value}
			processTemplate(fileName, value+".go", newData)
		}
	}

}
