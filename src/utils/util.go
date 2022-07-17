package utils

import (
	"fmt"
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

	"github.com/andodeki/gen_project/builder/codegen_builder/src/data"
	"github.com/andodeki/gen_project/builder/codegen_builder/src/process"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GenerateUtils(projectDefaults data.Data) {
	dataUtils := data.Data{
		ProductTargets: data.ProductTarget{
			OutputFilename: []string{"logger", "errors", "ulid", "helpers", "context"},
			DSource: data.Datasource{
				DBS: []string{"logger", "errors", "ulid", "helpers", "context"},
			},
		},
		ProductSource:   projectDefaults.ProductSource,
		Company:         projectDefaults.Company,
		ProductFileType: ".go",
		FolderPathName:  projectDefaults.ProjectFolder + "/src/util/",
		ConcreteTargets: []string{"logger", "errors", "ulid", "helpers", "context"},
	}
	// dataErrors:= data.Data{
	// 	ProductTargets: data.ProductTarget{
	// 		OutputFilename: []string{"logger", "errors"},
	// 		DSource: data.Datasource{
	// 			DBS: []string{"logger", "errors"},
	// 		},
	// 	},
	// 	ProductSource:   "propertylisting",
	// 	ProductFileType: ".go",
	// 	FolderPathName:  "./tmp/src/util/",
	// 	ConcreteTargets: []string{"logger", "errors"},
	// }
	// process.ProcessConcreteTargets("src/utils/logger.tmpl", dataLogger)
	// process.ProcessConcreteTargets("src/utils/errors.tmpl", dataErrors)

	process.ProcessConcreteTargets("src/utils/utils.tmpl", dataUtils)

}

func ReadFolder(filename string) []fs.FileInfo {
	files, err := ioutil.ReadDir(filename)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func ToOutPutDataType(ty string, op data.Data) interface{} {
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
		case "JSON":
			return "*string"
		case "TEXT":
			return "*string"
		case "TIMESTAMPTZ":
			return "*time.Time"
		case "bytea":
			return "*[]byte "
		case "BOOLEAN":
			return "*bool"
		case "BIGINT":
			return "*int64"
		case "INTEGER":
			return "*int64"
		case "FLOAT":
			return "*float64"
		default:
			// fmt.Printf("ty: %v\n", ty)FLOAT
			new_ty := strings.Replace(
				cases.Title(
					language.Und, cases.NoLower,
				).String(
					strings.Replace(ty, "_", " ", -1),
				), " ", "", -1,
			)
			// fmt.Printf("new_ty: %v\n", new_ty)
			return new_ty
		}
		// if ty == "UUID"{}
	} else {
		return ty
	}

	return ""
}

func RemoveLinesContainingAny(input string, toRemove []string) string {
	if !strings.HasSuffix(input, "\n") {
		input += "\n"
	}

	lines := strings.Split(input, "\n")

	for i, line := range lines {
		for _, rm := range toRemove {
			if strings.Contains(line, rm) {
				lines = append(lines[:i], lines[i+1:]...)
			}
		}
	}

	input = strings.Join(lines, "\n")
	input = strings.TrimSpace(input)
	input += "\n"

	return input
}

func RegxReplaceAllString(v string, charx string) string {
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

func TableFieldsFunc(jpos string, pos string, name string, tableFields []string, size int) []string {

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
	lines = Delete_empty(lines)
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
func Delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

// betweenSpecial
func BetweenSpecial(value string, a string, b string, size int) string {
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
func Before(value string, a string) string {
	// Get substring before a string.
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}
func After(value string, a string) string {
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

// https://www.dotnetperls.com/between-before-after-go
func Between(value string, a string, b string) string {
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

func typeOfVar(s interface{}) string {
	return fmt.Sprint(reflect.TypeOf(s)) //reflect.TypeOf(s)
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

func readFolder(filename string) []fs.FileInfo {
	files, err := ioutil.ReadDir(filename)
	if err != nil {
		log.Fatal(err)
	}
	return files
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

func lt() interface{} {
	return template.HTML(fmt.Sprint(`<`))
}
