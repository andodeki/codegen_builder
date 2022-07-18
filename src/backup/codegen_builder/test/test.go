package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"log"
	"os"

	"github.com/Masterminds/sprig"
)

type Data struct {
	ProductTargets  ProductTarget
	ConcreteTargets []string
	// Properties      []Property
}

type ProductTarget struct {
	Name    string
	Structs map[TableName][]Property
}

type TableName struct {
	Name          string
	ProductSource string
}

type Property struct {
	Name     string
	TypeName string
}

func main() {
	pt := new(ProductTarget)
	pt.Structs = make(map[TableName][]Property)
	tblname := TableName{"house", "propertylist"}
	tblname2 := TableName{"toilet", "propertylist"}

	pt.Structs[tblname] = []Property{{"windowType", "string"}, {"doorType", "string"}, {"floor", "int"}}
	pt.Structs[tblname2] = []Property{{"windowType", "string"}, {"doorType", "string"}, {"floor", "int"}}

	// keys := make([]string, 0, len(pt.Structs))
	// for k := range pt.Structs {
	// 	keys = append(keys, k)
	// }
	data := Data{
		ProductTargets: ProductTarget{
			Name:    "house",
			Structs: pt.Structs,
		},
		ConcreteTargets: []string{"normal", "igloo"},
		// Properties:      []Property{{"windowType", "string"}, {"doorType", "string"}, {"floor", "int"}},
	}
	fmt.Printf("data: %q\n", data)
	// processTemplate("iBuilder.tmpl", "iBuilder.go", data)
	// for _, v := range data.ProductTargets {
	// 	processTemplate("ProductTarget.tmpl", v.Name+".go", data)

	// }
	processTemplate("ProductTarget.tmpl", data.ProductTargets.Name+".go", data)

	// processTemplate("director.tmpl", "director.go", data)
	// processTemplate("sample_main.tmpl", "sample_main.go", data)
	// processConcreteTargets("ConcreteTarget.tmpl", data)
	fmt.Println("Remember to edit the files that contain the Concrete Targets!")
}

func processTemplate(fileName string, outputFile string, data Data) {
	tmpl := template.Must(template.New("").Funcs(sprig.FuncMap()).ParseFiles(fileName))
	var processed bytes.Buffer
	err := tmpl.ExecuteTemplate(&processed, fileName, data)
	if err != nil {
		log.Fatalf("Unable to parse data into template: %v\n", err)
	}
	formatted, err := format.Source(processed.Bytes())
	if err != nil {
		log.Fatalf("Could not format processed template: %v\n", err)
	}
	// formatted := processed.Bytes()
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
