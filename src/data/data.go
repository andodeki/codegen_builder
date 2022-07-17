package data

import "io/fs"

type Data struct {
	ProductTargets  ProductTarget
	ProductSource   string
	Company         string
	ProjectFolder   string
	ProductFileType string
	Workfile        string
	FolderPathName  string
	ConcreteTargets []string
	Properties      []Property
	Enums           []Enum
	DBIndexes       []DBIndex
}

type ProductTarget struct {
	OutputFilename      []string
	ProductFileTypeFrom string
	Name                string
	Structs             map[TableName][]Property
	Types               map[string][]Enum
	Index               map[IndexName][]DBIndex
	Indexes             []map[string][]string
	YAML                []map[string]interface{}
	Config              []map[string]interface{}
	// DBSMap              map[string][]string
	// DBSMapFields        map[string][]string
	// DBS                 []string
	DSource Datasource

	/// To Edit Below structs fields
	StructNames   []TableName
	TableNames    []TableName
	TypeNames     []string
	IndexNames    []IndexName
	ProductSource string
}
type Datasource struct {
	DBSMap       map[string][]string
	DBSMapFields map[string][]string
	DBS          []string
	IMethods     map[string][][]string
}

type TableName struct {
	Name          string
	ProductSource string
}

type IndexName struct {
	Name  string
	Table string
}

// var tableName, typeName, indexName string

// var vals = []string{"TABLE ", "TYPE ", "INDEX "}

// var jsonString [][]string

// var ExtractedData2 = struct {
// 	TableName, TypeName, TndexName string
// 	Vals                           []string
// 	JsonString                     [][]string
// }{"", "", "",
// 	[]string{"TABLE ", "TYPE ", "INDEX "},
// 	[][]string{{}},
// }

type ExtractedData struct {
	TableName, TypeName, IndexName, Wrkfile string
	Vals                                    []string
	JsonString                              [][]string
	WrkFolder                               []fs.FileInfo
}

type Property struct {
	Name     string
	TypeName interface{}
	Tag      string
}

type Enum struct {
	Name       string
	CustomType string
}

type DBIndex struct {
	Name      string
	TableName string
}

func (prop *Data) AddProperty(p Property) []Property {
	prop.Properties = append(prop.Properties, p)
	return prop.Properties
}
func (prop *Data) AddProductTargetProperty(p ProductTarget) map[TableName][]Property {
	prop.ProductTargets.Structs = p.Structs //append(prop.Properties, p)
	return prop.ProductTargets.Structs
}
func (prop *Data) ClearProductTargetProperty() {
	// for k := range prop.ProductTargets.Structs {
	// 	delete(prop.ProductTargets.Structs, k)
	// }

	// p := new(ProductTarget)
	// prop.ProductTargets.Structs = p.Structs
	prop.ProductTargets.Structs = nil

	// return prop.ProductTargets.Structs
}
func (prop *Data) AddTable(p Property) []Property {
	prop.Properties = append(prop.Properties, p)
	return prop.Properties
}
func (e *Data) AddType(p Enum) []Enum {
	e.Enums = append(e.Enums, p)
	return e.Enums
}

func (prop *Data) AddProductTargetTypes(p ProductTarget) map[string][]Enum {
	// prop.ProductTargets.Types = append(prop.ProductTargets.Types, p)
	prop.ProductTargets.Types = p.Types
	return prop.ProductTargets.Types
}

func (t *Data) ClearProductTargetTypes() {
	t.ProductTargets.Types = nil
}
func (prop *Data) AddProductTargetIndex(p ProductTarget) map[IndexName][]DBIndex {
	prop.ProductTargets.Index = p.Index //append(prop.ProductTargets.Indexes, p)
	return prop.ProductTargets.Index
}

func (t *Data) ClearProductTargetIndex() {
	t.ProductTargets.Indexes = nil
}
func (prop *Data) ClearProperty() []Property {
	prop.Properties = nil
	return prop.Properties
}

func (prop *Data) ClearProductTarget() []TableName {
	prop.ProductTargets.TableNames = nil
	return prop.ProductTargets.TableNames
}

func (e *Data) ClearEnum() []Enum {
	e.Enums = nil
	return e.Enums
}
func (t *Data) ClearTypes() []string {
	t.ProductTargets.TypeNames = nil
	return t.ProductTargets.TypeNames
}

func (i *Data) AddDBIndex(p DBIndex) []DBIndex {
	i.DBIndexes = append(i.DBIndexes, p)
	return i.DBIndexes
}
func (i *Data) ClearDBIndex() []DBIndex {
	i.DBIndexes = nil
	return i.DBIndexes
}

type appConfig struct {
	Version  string
	IsDev    bool
	Features appFeatures
}

type appFeatures struct {
	HttpServer      map[string]interface{}
	PostgresClients []map[string]interface{}
	MonetDBClients  string
	RedisDBClients  string
	ScyllaDBClients string
	JaegerClients   string
	Queue           string
	Email           string
	PasetoToken     string
	LogConfig       bool
}
