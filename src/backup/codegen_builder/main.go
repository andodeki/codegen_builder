package main

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
	"strings"

	"github.com/Masterminds/sprig"
)

type Data struct {
	ProductTarget   string
	ConcreteTargets []string
	Properties      []Property
}

type Property struct {
	Name     string
	TypeName string
}

func main() {
	data := Data{
		ProductTarget:   "house",
		ConcreteTargets: []string{"normal", "igloo"},
		Properties: []Property{
			{"windowType", "string"},
			{"doorType", "string"},
			{"floor", "int"},
		},
	}
	// data := Data{
	// 	ProductTarget:   "datasources",
	// 	ConcreteTargets: []string{"postgresdb", "scylladb"},
	// 	Properties: []Property{
	// 		{"windowType", "string"},
	// 		{"doorType", "string"},
	// 		{"floor", "int"},
	// 	},
	// }
	// processTemplate("iBuilder.tmpl", "iBuilder.go", data)
	processTemplate("ProductTarget.tmpl", data.ProductTarget+".go", data)
	// processTemplate("ProductTarget.tmpl", data.ProductTarget+".go", data)
	// processTemplate("director.tmpl", "director.go", data)
	// processTemplate("sample_main.tmpl", "sample_main.go", data)
	// processConcreteTargets("ConcreteTarget.tmpl", data)
	fmt.Println("Remember to edit the files that contain the Concrete Targets!")
}

func processTemplate(fileName string, outputFile string, data Data) {
	// tmpl := template.Must(template.New("").Funcs(sprig.FuncMap()).ParseFiles(fileName))
	// tmpl := template.Must(template.New("").Funcs(sprig.FuncMap()).ParseGlob("templates/Products/" + fileName))
	tmpl, errp := findAndParseTemplates(".", sprig.FuncMap())
	// // tmpl, errp := parseTemplate()
	if errp != nil {
		log.Print(errp)
	}

	var processed bytes.Buffer
	err := tmpl.ExecuteTemplate(&processed, fileName, data)
	if err != nil {
		log.Fatalf("Unable to parse data into template: %v\n", err)
	}
	formatted, err := format.Source(processed.Bytes())
	if err != nil {
		log.Fatalf("Could not format processed template: %v\n", err)
	}
	outputPath := "./tmp/" + outputFile
	fmt.Println("Writing file: ", outputPath)
	f, _ := os.Create(outputPath)
	w := bufio.NewWriter(f)
	w.WriteString(string(formatted))
	w.Flush()
}

func processConcreteTargets(fileName string, data Data) {
	for _, value := range data.ConcreteTargets {
		newData := data
		newData.ConcreteTargets = []string{value}
		processTemplate(fileName, value+".go", newData)
	}
}
func parseTemplate() (*template.Template, error) {
	templateBuilder := template.New("").Funcs(sprig.FuncMap())
	if t, _ := templateBuilder.ParseGlob("/*/*/*/*/*.tmpl"); t != nil {
		templateBuilder = t
	}
	if t, _ := templateBuilder.ParseGlob("/*/*/*/*.tmpl"); t != nil {
		templateBuilder = t
	}
	if t, _ := templateBuilder.ParseGlob("/*/*/*.tmpl"); t != nil {
		templateBuilder = t
	}
	if t, _ := templateBuilder.ParseGlob("/*/*.tmpl"); t != nil {
		templateBuilder = t
	}
	return templateBuilder.ParseGlob("/*.tmpl")

}
func findAndParseTemplates(rootDir string, funcMap template.FuncMap) (*template.Template, error) {
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
			name := path[pfx:]
			t := root.New(name).Funcs(funcMap)
			_, e2 = t.Parse(string(b))
			if e2 != nil {
				return e2
			}
		}
		return nil
	})
	return root, err

}
