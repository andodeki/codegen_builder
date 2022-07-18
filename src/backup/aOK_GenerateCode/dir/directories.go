package dir

func PopulatedDirectoriesSlice() (Directories []string) {
	Directories = append(SrcDir, DatasourcesDir...)
	// Directories = append(SrcDir, UtilDir...)
	Directories = append(SrcDir, CacheDir...)
	// Directories = append(Directories, QueueFiles...)
	// Directories = append(Directories, ScriptsFiles...)
	// Directories = append(Directories, ConfigFiles...)

	return Directories
}

var SrcDir = []string{
	"configuration",
	"webclient",
	"src/util",
	"src/models",
	"src/auth",
	"src/queue",
	"src/email",
	"src/keys",
	"src/scripts",
	"src/util",
	"src/datasources",
	"src/repository",
	"src/services",
	"src/apiclient",
	"src/api",
	"src/config",
	"src/startup",
	"src/httpServer",
}

var DatasourcesDir = []string{
	"src/datasources/migrations",
	"src/datasources/migrations/cqlmigrations",
}

// var UtilDir = []string{
// 	"src/util",
// }

// var UtilDir = []string{
// 	"src/util",
// }

var CacheDir = []string{
	"src/cache",
	"src/cache/redis",
	"src/cache",
	"src/cache/memcached",
}

// // var folders = []string{}
// var Directories = []string{
// 	"configuration",
// 	"src/models",
// 	"src/auth",
// 	"src/queue",
// 	"src/email",
// 	"src/keys",
// 	"src/scripts",
// 	"src/util",
// 	"src/datasources",
// 	"src/repository",
// 	"src/services",
// 	"src/apiclient",
// 	"src/api",
// 	"src/config",
// 	"src/startup",
// 	"src/httpServer",
// 	"src/cache",
// 	"webclient",
// }
