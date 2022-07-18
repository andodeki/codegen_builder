package dir

func PopulatedFileSlice() (Files []string) {
	Files = append(SrcFiles, DatasourcesFiles...)
	Files = append(Files, QueueFiles...)
	Files = append(Files, ScriptsFiles...)
	Files = append(Files, ConfigFiles...)
	Files = append(Files, HttpServerFiles...)
	Files = append(Files, UtilFiles...)

	return Files
}

var SrcFiles = []string{
	"main.go",
	"configuration.go",
	"startup.go",
	"postgresdb.go",
	"scylladb.go",
	"httpServer.go",
	"redisdb.go",
	"monetdb.go",
	"pprof.go",
}
var EmailFiles = []string{
	"queue/email.go",
	"queue/fetchEmail.go",
}
var UtilFiles = []string{
	"util/logger.go",
	"queue/error.go",
}
var ModelsFiles = []string{
	"models/base/user.go",
	"models/base/session.go",
	"models/base/session.go",
	"models/base/session.go",
	"models/base/session.go",
}

var ScriptsFiles = []string{
	"scripts/gomod.sh",
}
var QueueFiles = []string{
	"queue/email_queue.go",
}
var DatasourcesFiles = []string{
	"datasources/redis.go",
	"datasources/monetDBClient.go",
	"datasources/postgresClient.go",
	"datasources/scylla.go",
	"datasources/transaction.go",
	"datasources/migrate.go",
}

var HttpServerFiles = []string{
	"httpServer/httpServer.go",
	"httpServer/layer_test.go",
	"httpServer/layers.go",
	"httpServer/postgresEnabled.go",
	"httpServer/scyllaDBEnabled.go",
	"httpServer/serverControllers.go",
	"httpServer/serverServices.go",
}
var ConfigFiles = []string{
	"config/config	.go",
	"config/getters.go",
	"config/marshaler.go",
	"config/types.go",
	"config/config_file.go",
	"config/config_test.go",
}
