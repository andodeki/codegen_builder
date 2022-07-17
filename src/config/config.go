package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/andodeki/gen_project/builder/codegen_builder/src/data"
	"github.com/andodeki/gen_project/builder/codegen_builder/src/process"
)

func ConfigGenerator(dataYAML data.Data, projectDefaults data.Data) {
	cfg := new(data.ProductTarget)
	cfg.Config = dataYAML.ProductTargets.YAML
	dataConfig := data.Data{
		ProductTargets: data.ProductTarget{
			OutputFilename:      []string{"types", "getters", "interfaces", "marshaler", "config"},
			Config:              cfg.Config,
			ProductFileTypeFrom: dataYAML.ProductFileType,
		},
		ProductSource:   projectDefaults.ProductSource,
		Company:         projectDefaults.Company,
		ProductFileType: ".go",
		FolderPathName:  projectDefaults.ProjectFolder + "/src/config/",
		ConcreteTargets: []string{},
	}

	for _, v := range dataConfig.ProductTargets.OutputFilename {
		if v == "types" {
			process.ProcessNonConcreteTargets("src/config/typesconfig.tmpl", dataConfig, v)

		} else if v == "getters" {
			process.ProcessNonConcreteTargets("src/config/gettersconfig.tmpl", dataConfig, v)

		} else if v == "interfaces" {
			process.ProcessNonConcreteTargets("src/config/interfacesconfig.tmpl", dataConfig, v)

		} else if v == "marshaler" {
			process.ProcessNonConcreteTargets("src/config/marshalerconfig.tmpl", dataConfig, v)

		} else if v == "config" {
			process.ProcessNonConcreteTargets("src/config/appconfig.tmpl", dataConfig, v)

		}

	}
}

func YAMLGenerator(projectDefaults data.Data) (data.Data, map[string][]string, []string, []string) {
	dataYAML, name, kvdbs, dbs, confServices := generateYAML(projectDefaults)

	for _, v := range dataYAML.ProductTargets.OutputFilename {
		if v == "config" {
			process.ProcessNonConcreteTargets("src/config/config.tmpl", dataYAML, name)
		}
	}
	return dataYAML, kvdbs, dbs, confServices
}

func generateYAML(projectDefaults data.Data) (data.Data, string, map[string][]string, []string, []string) {
	yml := new(data.ProductTarget)
	yml.YAML = make([]map[string]interface{}, 0)
	value := map[string]interface{}{
		"Version": "0",
		"IsDev":   "True",
		"Configuration": []map[string]interface{}{
			{"Enabled": "True"},
			{"Production": "prod"},
			{"Development": "dev"},
			{"Testing": "test"},
		},
		"HttpServer": []map[string]interface{}{
			{"Enabled": "True"},
			{"ServerName": "api.propertylist.com"},
			{"Bind": "0.0.0.0"},
			{"Port": 8089},
			{"LogRequests": "True"},
			{"Cookie": "True"},
		},
		"PostgresDBClients": []map[string][]map[string]interface{}{
			{
				"0-postgres": []map[string]interface{}{
					{"Enabled": "True"},
					{"DatabaseHost": "10.38.195.235"},
					{"DatabasePort": 5432},
					{"DatabaseUsername": "propertylistingadmin"},
					{"DatabasePassword": "password"},
					{"DatabaseDriver": "postgres"},
					{"DatabaseName": "propertylisting"},
					{"DatabaseTimeout": 5000},
					{"DatabaseSSLMode": "disable"},
					{"MigrationDir": "./datasources"},
					{"DownMigration": "False"},
				},
			},
		},
		"ScyllaDBClients": []map[string][]map[string]interface{}{
			{
				"0-scylla": []map[string]interface{}{
					{"Enabled": "True"},
					{"Name": "scylla-node0"},
					{"ScyllaHosts": []string{"10.38.195.126", "10.38.195.20", "10.38.195.158"}},
					{"Username": "propertylistingadmin"},
					{"Password": "password"},
					{"DatabaseTimeout": 5000},
					{"Keyspace": "propertylisting"},
					{"Class": "NetworkTopologyStrategy"},
					{"ReplicationFactor": 3},
					{"MigrationDir": "./datasources/migrations/cqlmigrations"},
					{"DurableWrites": "True"},
				},
			},
			{
				"1-scylla": []map[string]interface{}{
					{"Enabled": "True"},
					{"Name": "scylla-node0"},
					{"ScyllaHosts": []string{"10.38.195.20"}},
					{"Username": "propertylistingadmin"},
					{"Password": "password"},
					{"Keyspace": "propertylisting"},
					{"Class": "NetworkTopologyStrategy"},
					{"ReplicationFactor": 3},
					{"MigrationDir": "./datasources/migrations/cqlmigrations"},
					{"DurableWrites": "True"},
				},
			},
			{
				"2-scylla": []map[string]interface{}{
					{"Enabled": "True"},
					{"Name": "scylla-node0"},
					{"ScyllaHosts": []string{"10.38.195.158"}},
					{"Username": "propertylistingadmin"},
					{"Password": "password"},
					{"Keyspace": "propertylisting"},
					{"Class": "NetworkTopologyStrategy"},
					{"ReplicationFactor": 3},
					{"MigrationDir": "./datasources/migrations/cqlmigrations"},
					{"DurableWrites": "True"},
				},
			},
		},
		"MonetDBClients": []map[string][]map[string]interface{}{
			{
				"0-monetdb": []map[string]interface{}{
					{"Enabled": "False"},
					{"DatabaseHost": "10.9.49.118"},
					{"DatabasePort": 50000},
					{"DatabaseUsername": "propertylistingadmin"},
					{"DatabasePassword": "password"},
					{"DatabaseTimeout": 5000},
					{"DatabaseDriver": "monetdb"},
					{"DatabaseName": "propertylisting"},
					{"DatabaseSSLMode": "disable"},
					{"MigrationDir": "./datasources"},
					{"DownMigration": "False"},
				},
			},
		},
		"RedisDBClients": []map[string][]map[string]interface{}{
			{
				"0-redis": []map[string]interface{}{
					{"Enabled": "True"},
					{"Name": "propertylistingapi"},
					{"RedisPassword": ""},
					{"RedisAddress": "10.38.195.64:6379"},
					{"DatabaseUsername": "propertylistingadmin"},
					{"DatabaseTimeout": 5000},
					{"RedisDatabase": 0},
					{"LogMessages": "True"},
				},
			},
		},
		"JaegerClients": []map[string][]map[string]interface{}{
			{
				"0-jaeger": []map[string]interface{}{
					{"Enabled": "False"},
					{"JaegerServiceName": "propertylist_jaeger"},
					{"JaegerEndpoint": "http://localhost:14268/api/traces"},
					{"LogMessages": "True"},
				},
			},
		},
		"Email": []map[string]interface{}{
			{"Enabled": "True"},
			{"Host": "smtp.gmail.com"},
			{"Username": "andrewodeki@gmail.com"},
			{"Password": "ICU@i4cu@ICU"},
			{"InsecureSkipVerify": "True"},
			{"RecipientUsername": "andodeki@gmail.com"},
			{"RecipientPassword": "icuarchicad17i4cuicu"},
			{"RecipientHost": "smtp.gmail.com:993"},
			{"Subject": "System Admin Password"},
			{"LogMessages": "True"},
		},
		"Queue": []map[string]interface{}{
			{"Enabled": "True"},
			{"MaxWorkers": 10},
			{"LogMessages": "True"},
		},
		"PasetoToken": []map[string]interface{}{
			{"Enabled": "True"},
			{"TokenSymmetricKey": "12345678901234567890123456789012"},
			{"AccessTokenDuration": 30},
			{"RefreshTokenDuration": 30},
			{"AccessTokenDurationDev": 30},
			{"RefreshTokenDurationDev": 60},
			{"IsDev": "True"},
		},
	}

	yml.YAML = append(yml.YAML, value)

	dataYAML := data.Data{
		ProductTargets: data.ProductTarget{
			OutputFilename: []string{"config"},
			YAML:           yml.YAML,
		},
		ProductSource:   projectDefaults.ProductSource,
		Company:         projectDefaults.Company,
		ProductFileType: ".yaml",
		FolderPathName:  projectDefaults.ProjectFolder + "/configurations/",
		ConcreteTargets: []string{},
	}

	kvdbs := make(map[string][]string)
	kvconf := make(map[string][]string)
	dbs := make([]string, 0)
	conf := make([]string, 0)
	for ks, vs := range value {
		//OUTPUT: [RedisDBClients PostgresDBClients ScyllaDBClients MonetDBClients]
		if fmt.Sprint(reflect.TypeOf(vs)) == "[]map[string][]map[string]interface {}" && strings.Contains(ks, "DB") {
			if rec, ok := vs.([]map[string][]map[string]interface{}); ok {
				var list []string
				for key, val := range rec[0] {
					_ = key
					_ = val
					for klist, vlist := range val {
						_ = vlist
						_ = klist
						for kmap := range vlist {
							list = append(list, kmap)
						}
					}
				}
				kvdbs[ks] = list

			} else {
				fmt.Printf("record not a map[string]interface{}: %v\n", reflect.TypeOf(vs))
			}
		}

		//OUTPUT: [RedisDBClients JaegerClients PostgresDBClients ScyllaDBClients MonetDBClients]
		if fmt.Sprint(reflect.TypeOf(vs)) == "[]map[string][]map[string]interface {}" {
			if rec, ok := vs.([]map[string][]map[string]interface{}); ok {
				var list []string
				for key, val := range rec[0] {
					_ = key
					_ = val
					for klist, vlist := range val {
						_ = vlist
						_ = klist
						for kmap := range vlist {
							list = append(list, kmap)
						}
					}
				}
				kvconf[ks] = list
			}
		} else if fmt.Sprint(reflect.TypeOf(vs)) == "[]map[string]interface {}" {
			if rec, ok := vs.([]map[string]interface{}); ok {
				var list []string
				for key, val := range rec[0] {
					_ = key
					_ = val
				}
				kvconf[ks] = list
			}
			// conf = append(conf, ks)

		} else {
			// conf = append(conf, ks)

		}
	}

	for k := range kvdbs {
		dbs = append(dbs, k)
	}
	for k := range kvconf {
		conf = append(conf, k)
	}
	// fmt.Printf("dbs: %v\n", conf)

	return dataYAML, "config", kvdbs, dbs, conf
}
