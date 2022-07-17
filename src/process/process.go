package process

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/Masterminds/sprig"
	"github.com/andodeki/gen_project/builder/codegen_builder/src/data"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ProcessNonConcreteTargets(fileName string, data data.Data, name string) {
	if data.ProductFileType == ".cql" {
		newData := data
		newData.ProductTargets = data.ProductTargets
		path := data.FolderPathName
		// fmt.Printf("path: %v\n", path)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			if err := EnsureDir(path); err != nil {
				log.Println("Directory creation failed with error: " + err.Error())
			}
			if cherr := ChownR(path, 1000, 1000); cherr != nil {
				log.Printf("cannot chmod to file 1:%v", cherr)
			}
		}
		// name := data.ProductTargets.OutputFilename
		// for _, v := range name {
		processTemplate(fileName, path+name+".cql", newData)
	}
	if data.ProductFileType == ".go" {
		newData := data
		newData.ProductTargets = data.ProductTargets
		newData.ProductTargets.DSource.DBS = data.ProductTargets.DSource.DBS
		path := data.FolderPathName
		if _, err := os.Stat(path); os.IsNotExist(err) {
			if err := EnsureDir(path); err != nil {
				log.Println("Directory creation failed with error: " + err.Error())
			}
			if cherr := ChownR(path, 1000, 1000); cherr != nil {
				log.Printf("cannot chmod to file 2:%v", cherr)
			}
		}
		// name := data.ProductTargets.OutputFilename
		// for _, v := range name {
		// processTemplate(fileName, path+v+".go", newData)

		// }
		processTemplate(fileName, path+name+".go", newData)
	}
	if data.ProductFileType == ".yaml" {
		// fmt.Printf("data.ProductTargets: %v\n", data.ProductTargets)
		newData := data
		newData.ProductTargets = data.ProductTargets
		path := data.FolderPathName
		if _, err := os.Stat(path); os.IsNotExist(err) {
			if err := EnsureDir(path); err != nil {
				log.Println("Directory creation failed with error: " + err.Error())
			}
			if cherr := ChownR(path, 1000, 1000); cherr != nil {
				log.Printf("cannot chmod to file 3:%v", cherr)
			}
		}
		// name := data.ProductTargets.OutputFilename
		// for _, v := range name {
		// processTemplate(fileName, path+v+".yaml", newData)

		// }
		processTemplate(fileName, path+name+".yaml", newData)
	}

}

func processTemplate(fileName string, outputFile string, data data.Data) {
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
		"toJS":                     toJS,
		"lthan":                    lt,
		"replace":                  replace,
		"displaySpecialCharacters": displaySpecialCharacters,
		"contains":                 contains,
		"typeOfVar":                typeOfVar,
		"TitleRemoveUnderscore":    TitleRemoveUnderscore,
		"typeOfVarSpecial":         typeOfVarSpecial,
		"stringToInt":              stringToInt,
		"IsDigit":                  IsDigit,
		"replaceNonAlpa":           regxReplaceAllString,
		"toOutPutDataTypeYaml":     toOutPutDataTypeYaml,
		"lowerTitle":               MakeFirstLowerCase,
		"safeURL":                  func(u string) template.URL { return template.URL(u) },
	}
	// t.Funcs(template.FuncMap{
	// 	"stringSlice": func(s string, i, j int) string {
	// 		return s[i:j]
	// 	}
	// })typeOfVarSpecial
	// tmpl := template.Must(template.New("").Funcs(sprig.FuncMap()).Funcs(helpers).ParseFiles(fileName))
	tmpl, errp := findAndParseTemplates(".", sprig.FuncMap(), helpers)
	if errp != nil {
		log.Print(errp)
	}
	var processed bytes.Buffer
	err := tmpl.ExecuteTemplate(&processed, fileName, data)
	if err != nil {
		log.Fatalf("Unable to parse data into template: %v\n", err)
	}

	if data.ProductFileType == ".cql" {
		// sprocessed := processed.String()
		formatted := processed.Bytes()
		outputPath := outputFile
		// fmt.Printf("outputPath: %v\n", outputPath)
		fmt.Println("Writing file: ", outputPath)
		f, _ := os.Create(outputPath)
		w := bufio.NewWriter(f)
		newlineString := strings.Replace(string(formatted), "<br>", "\n", -1)
		tabString := strings.Replace(newlineString, "<brt>", "\t", -1)
		w.WriteString(tabString)
		w.Flush()
	} else if data.ProductFileType == ".go" {
		// fmt.Printf("fileName: %v\n", fileName)
		formatted, err := format.Source(processed.Bytes())
		if err != nil {
			log.Fatalf("Could not format processed template: %v\n", err)
		}
		// formatted := processed.Bytes()
		outputPath := outputFile
		fmt.Println("Writing file: ", outputPath)
		f, createErr := os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if createErr != nil {
			log.Printf("cannot create file:%v", createErr)
		}
		w := bufio.NewWriter(f)
		newlineString := strings.Replace(string(formatted), "<br>", "\n", -1)
		tabString := strings.Replace(newlineString, "<brt>", "\t", -1)
		rmQuotes := strings.Replace(tabString, "&#34;", "\"", -1)
		w.WriteString(rmQuotes)
		w.Flush()
	} else if data.ProductFileType == ".yaml" {
		formatted := processed.Bytes()
		outputPath := outputFile
		// fmt.Printf("outputPath: %v\n", outputPath)
		fmt.Println("Writing file: ", outputPath)
		f, _ := os.Create(outputPath)
		w := bufio.NewWriter(f)
		w.WriteString(string(formatted))
		w.Flush()
	}

}

func ProcessConcreteTargets(fileName string, data data.Data) {
	for _, value := range data.ProductTargets.DSource.DBS {
		// if strings.Contains(value, "Postgres") {
		// fmt.Printf("value: %v\n", value)
		// }
		for _, v := range data.ConcreteTargets {
			if strings.Contains(strings.ToLower(value), v) {
				newData := data
				newData.ConcreteTargets = []string{v}
				newData.ProductTargets.DSource.DBS = []string{value}
				path := data.FolderPathName
				if _, err := os.Stat(path); os.IsNotExist(err) {
					if err := EnsureDir(path); err != nil {
						log.Println("Directory creation failed with error: " + err.Error())
					}
					if cherr := ChownR(path, 1000, 1000); cherr != nil {
						log.Printf("cannot chmod to file 2:%v", cherr)
					}
				}
				processTemplate(fileName, path+v+".go", newData)
			} else if strings.Contains(strings.ToLower(value), v) && strings.Contains(fileName, v) {
				newData := data
				newData.ConcreteTargets = []string{v}
				newData.ProductTargets.DSource.DBS = []string{value}
				path := data.FolderPathName
				if _, err := os.Stat(path); os.IsNotExist(err) {
					if err := EnsureDir(path); err != nil {
						log.Println("Directory creation failed with error: " + err.Error())
					}
					if cherr := ChownR(path, 1000, 1000); cherr != nil {
						log.Printf("cannot chmod to file 2:%v", cherr)
					}
				}
				processTemplate(fileName, path+v+".go", newData)
			}
		}

	}
	// }
	// for _, value := range data.ConcreteTargets {
	// 	newData := data
	// 	newData.ConcreteTargets = []string{value}
	// 	path := data.FolderPathName
	// 	if _, err := os.Stat(path); os.IsNotExist(err) {
	// 		if err := EnsureDir(path); err != nil {
	// 			log.Println("Directory creation failed with error: " + err.Error())
	// 		}
	// 		if cherr := ChownR(path, 1000, 1000); cherr != nil {
	// 			log.Printf("cannot chmod to file 2:%v", cherr)
	// 		}
	// 	}
	// 	processTemplate(fileName, path+value+".go", newData)
	// }
}
func findAndParseTemplates(rootDir string, funcMap template.FuncMap, helpers template.FuncMap) (*template.Template, error) {
	cleanRoot := filepath.Clean(rootDir)
	pfx := len(cleanRoot) + 1
	root := template.New("")

	err := filepath.Walk(cleanRoot, func(path string, info fs.FileInfo, e1 error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".tmpl") {
			if e1 != nil {
				return e1
			}
			b, e2 := ioutil.ReadFile(path)
			if e2 != nil {
				return e2
			}
			_ = pfx
			name := path[:]
			// fmt.Printf("name: %v\n", name)
			t := root.New(name).Funcs(funcMap).Funcs(helpers)
			_, e2 = t.Parse(string(b))
			if e2 != nil {
				return e2
			}
		}
		return nil
	})
	return root, err

}

func TitleRemoveUnderscore(s string) string {
	var new_s = strings.Replace(
		cases.Title(
			language.Und, cases.NoLower,
		).String(
			strings.Replace(s, "_", " ", -1),
		), " ", "", -1,
	)
	return new_s
}
func IsDigit(s string) bool {
	a := []rune(s)
	if len(a) < 2 {
		isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
		b := strings.IndexFunc(s, isNotDigit) == -1
		return b
	}
	return true
}
func regxReplaceAllString(v string, charx string) string {
	if charx != "" {
		val := fmt.Sprintf("[^a-zA-Z0-9%s]+", charx)
		reg, err := regexp.Compile(val)
		if err != nil {
			log.Fatal(err)
		}
		return reg.ReplaceAllString(v, "")
	} else if charx == "" {
		reg, err := regexp.Compile("[^a-zA-Z0-9]+")
		if err != nil {
			log.Fatal(err)
		}
		return reg.ReplaceAllString(v, "")
	}

	return ""
}

func tableFieldsFunc(jpos string, pos string, name string, tableFields []string, size int) []string {

	lines := strings.Split(jpos, "\n")
	// fmt.Printf("lines: %v\n", lines)
	for i, line := range lines {
		if strings.Contains(line, "--") {
			lines[i] = ""
		} else if strings.Contains(line, pos) {
			// min, _ := MinMax(line, ";")
			// fmt.Printf("<<<name>>>: %v\n", name+"(")

			// lin := betweenSpecial(line, name+"(", ";", min)
			// fmt.Printf("lin-pos: %v\n", lin)
			// lines[i] = lin
			lines[i] = ""
		} else {
			lines[i] = strings.Join(strings.Fields(line), " ")
		}
	}
	// fmt.Printf("lines: %v\n", lines)
	lines = delete_empty(lines)
	// fmt.Printf("lines: %v\n", lines)
	str1 := strings.Join(lines, " ")
	str2 := strings.Split(str1, ",")
	for _, i := range str2 {
		str2 = strings.Split(strings.TrimPrefix(i, " "), " ")
		tableFields = append(tableFields, strings.Join(str2[:2], " "))
	}
	return tableFields
}
func MinMax(line, char string) (min int, max int) {
	var input []int
	for i, j := range strings.Split(line, "") {
		if strings.Contains(j, char) {
			input = append(input, i)
		}
	}
	if len(char) > 1 {
		min = len(char)
		max = len(char)
	} else if len(char) <= 1 {
		min = input[0]
		max = input[0]
	}
	for _, value := range input {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
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

// betweenSpecial
func betweenSpecial(value string, a string, b string, size int) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	var posLast int
	var nextChr string
	if posFirst+size <= len(value) {
		nextChr = value[posFirst : posFirst+size]
	} else if posFirst+size >= len(value) {
		nextChr = value[posFirst:size]
	}
	if strings.Contains(nextChr, b) { // b = "("
		posLast = strings.Index(nextChr, b)
		if posLast == -1 {
			return ""
		}
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posFirst+posLast {
		str := fmt.Sprintf("\nposFirst: %d\nposFirstAdjusted: %d\nposFirst+posLast: %d", posFirst, posFirstAdjusted, posFirst+posLast)
		return str //fmt.Sprintf("\nposFirst: %d\nposFirstAdjusted: %d\nposFirst+posLast: %d", posFirst, posFirstAdjusted, posFirst+posLast)
	}

	return value[posFirstAdjusted : posFirst+posLast]
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

func typeOfVarSpecial(s interface{}) string {
	var intVar int
	if _, err := strconv.ParseInt(s.(string), 10, 64); err == nil {
		intVar, _ = strconv.Atoi(s.(string))
		return fmt.Sprint(reflect.TypeOf(intVar))
	} else if err != nil {
		a := []rune(s.(string))
		if _, err := strconv.ParseInt(string(a[0]), 10, 64); err == nil {
			intVar, _ = strconv.Atoi(s.(string))
			return fmt.Sprint(reflect.TypeOf(intVar))
		}
		// return fmt.Sprint(reflect.TypeOf(intVar))
	}
	a := []rune(s.(string))
	_ = a

	return fmt.Sprint(reflect.TypeOf(s)) //reflect.TypeOf(s)

}
func typeOfVar(s interface{}) string {
	// var w bool
	// var u []string
	// var v int
	// var r []map[string]interface{}
	// var t []map[string][]map[string]interface{}

	// if fmt.Sprint(reflect.TypeOf(s)) == fmt.Sprint(reflect.TypeOf(w)) {
	// 	return fmt.Sprint(reflect.TypeOf(s))
	// }
	// if fmt.Sprint(reflect.TypeOf(s)) == fmt.Sprint(reflect.TypeOf(v)) {
	// 	return fmt.Sprint(reflect.TypeOf(s))
	// }

	// if fmt.Sprint(reflect.TypeOf(s)) == fmt.Sprint(reflect.TypeOf(u)) {
	// 	return fmt.Sprint(reflect.TypeOf(s))
	// }
	// if fmt.Sprint(reflect.TypeOf(s)) == fmt.Sprint(reflect.TypeOf(r)) {
	// 	return fmt.Sprint(reflect.TypeOf(s))
	// }
	// if fmt.Sprint(reflect.TypeOf(s)) == fmt.Sprint(reflect.TypeOf(t)) {
	// 	return fmt.Sprint(reflect.TypeOf(s))
	// }

	// if _, err := strconv.ParseInt(s.(string), 10, 64); err == nil {
	// 	intVar, _ := strconv.Atoi(s.(string))
	// 	return fmt.Sprint(reflect.TypeOf(intVar))
	// }
	return fmt.Sprint(reflect.TypeOf(s)) //reflect.TypeOf(s)

}

func before(value string, a string) string {
	// Get substring before a string.
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}

func lowerTitle(s string) string {
	// str := cases.Title(language.Und, cases.NoLower).String(s)
	str := cases.Title(language.Und, cases.NoLower).String(s)
	return str
}
func EnsureDir(fileName string) error {
	dirName := filepath.Dir(fileName)
	// if _, serr := os.Stat(dirName); serr != nil {
	merr := os.MkdirAll(dirName, os.ModePerm)
	if merr != nil || os.IsExist(merr) {
		return merr
	}
	// }
	if cherr := ChownR(dirName, 1000, 1000); cherr != nil {
		log.Printf("cannot chmod to file:%v", cherr)
	}
	return nil
}

func MakeFirstLowerCase(s string) string {
	// https://go.dev/play/p/tu9Cum4EmP
	a := []rune(s)
	a[0] = unicode.ToLower(a[0])
	s = string(a)
	return s
}
func ChownR(path string, uid, gid int) error {
	return filepath.Walk(path, func(name string, info os.FileInfo, err error) error {
		if err == nil {
			err = os.Chown(name, uid, gid)
		}
		return err
	})
}

func toJS(s string) template.URL {
	return template.URL(s)
}

func replace(input, from, to string) string {
	return strings.Replace(input, from, to, -1)
}
func contains(input, value string) bool {
	return strings.Contains(input, value)
}

func stringToInt(s string) interface{} {
	i, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Sprint(s)
		//log.Print(err)
	}
	return i
}
func toOutPutDataTypeYaml(ty, data string) interface{} {
	if data == ".yaml" {
		switch ty {
		case "True":
			boolValue, err := strconv.ParseBool(ty)
			if err != nil {
				log.Print(err)
			}
			return boolValue
		case "False":
			boolValue, err := strconv.ParseBool(ty)
			if err != nil {
				log.Print(err)
			}
			return boolValue
		}
	}
	return ""

}

func displaySpecialCharacters(s string) string {
	tt := template.HTML(s)
	return string(tt)
}
func toOutPutDataType(ty string, op data.Data) interface{} {
	if op.ProductFileType == ".cql" {
		switch ty {
		case "UUID":
			return "UUID"
		case "TEXT":
			return "TEXT"
		case "TIMESTAMPTZ":
			return "TIMESTAMP"
		case "BOOLEAN":
			return "BOOLEAN"
		case "bytea":
			return "TEXT"
		case "BIGINT":
			return "BIGINT"
		case "INTEGER":
			return "BIGINT"
		default:
			return template.HTML(fmt.Sprintf(`frozen<%s>`, ty))
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
		case "BOOLEAN":
			return "*boolean"
		case "BIGINT":
			return "*int64"
		case "INTEGER":
			return "*int64"
		default:
			return strings.Replace(
				cases.Title(
					language.Und, cases.NoLower,
				).String(
					strings.Replace(ty, "_", " ", -1),
				), " ", "", -1,
			)
		}
		// if ty == "UUID"{}
	}
	return ""
}

func lt() interface{} {
	return template.HTML(fmt.Sprint(`<`))
}
