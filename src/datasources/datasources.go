package datasources

import (
	"fmt"
	"strings"

	"github.com/andodeki/gen_project/builder/codegen_builder/src/data"
	"github.com/andodeki/gen_project/builder/codegen_builder/src/process"
)

func DataSourcesTargets(
	kvdbs map[string][]string, dbs []string, dataYAML data.Data,
	// propts []data.Property, enums []data.Enum, idxs []data.DBIndex,
	projectDefaults data.Data,
) {
	cts := make([]string, 0)
	for _, v := range dbs {
		cts = append(cts, strings.ToLower(strings.Replace(v, "Clients", "", -1)))
	}
	// dsrcsfields := make([]string, 0)
	dsrcs := new(data.ProductTarget)
	dsrcs.DSource.DBSMapFields = make(map[string][]string)
	for k := range kvdbs {
		if strings.Contains(strings.ToLower(k), "postgres") {
			structfields := []string{"logger *util.Logger",
				fmt.Sprintf("config *config.%v", k),
				"client *sqlx.DB",
			}
			dsrcs.DSource.DBSMapFields[k] = structfields
		}
		if strings.Contains(strings.ToLower(k), "scylla") {
			structfields := []string{"logger *util.Logger",
				fmt.Sprintf("config *config.%v", k),
				"clientSession  *gocql.Session",
				"clientx *gocqlx.Session",
			}
			dsrcs.DSource.DBSMapFields[k] = structfields
		}
		if strings.Contains(strings.ToLower(k), "redis") {
			structfields := []string{"logger *util.Logger",
				fmt.Sprintf("config *config.%v", k),
				"client   *redis.Client",
				"Cache  *cache.Cache",
			}
			dsrcs.DSource.DBSMapFields[k] = structfields
		}
		if strings.Contains(strings.ToLower(k), "monet") {
			structfields := []string{"logger *util.Logger",
				fmt.Sprintf("config *config.%v", k),
				"client *sqlx.DB",
			}

			dsrcs.DSource.DBSMapFields[k] = structfields
		}

	}
	iterfacemethods := map[string][][]string{
		"Run":         {{"ctx context.Context", "retries int", "logger *util.Logger", "config"}, {"client ", "err error"}},
		"DBClient":    {{""}, {"."}},
		"Health":      {{"ctx context.Context"}, {"error"}},
		"waitForDB":   {{"ctx context.Context"}, {"error"}},
		"migrateDb":   {{"ctx context.Context", "logger *util.Logger", "config"}, {"error"}},
		"Transaction": {{"ctx context.Context", "fn func(ctx context.Context) error"}, {"error"}},
	}
	_ = iterfacemethods
	// {{if (isLast $idx $lenitm)}}{{- else}}{{end}}

	dataSources := data.Data{
		ProductTargets: data.ProductTarget{
			OutputFilename: []string{"postgres", "scylla", "monetdb", "marshaler", "config"},
			DSource: data.Datasource{
				DBSMap:       kvdbs,
				DBSMapFields: dsrcs.DSource.DBSMapFields,
				DBS:          dbs,
				IMethods:     iterfacemethods,
			},
			ProductFileTypeFrom: dataYAML.ProductFileType,
		},
		ProductSource:   "propertylisting",
		ProductFileType: ".go",
		FolderPathName:  projectDefaults.ProjectFolder + "/src/datasources/",
		ConcreteTargets: cts,

		// Properties: propts,
		// Enums:      enums,
		// DBIndexes:  idxs,
	}
	process.ProcessConcreteTargets("src/datasources/datasources.tmpl", dataSources)
}
