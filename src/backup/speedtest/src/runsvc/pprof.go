package runsvc

//pprof//
//[Email HttpServer Queue PasetoToken JaegerClients Email PasetoToken ScyllaDBClients 
//HttpServer PostgresDBClients MonetDBClients RedisDBClients Queue]
import (
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/andodeki/propertylisting/util"
)

func RunCpuProfile(fileName string, logger *util.Logger) (started bool) {
	if fileName == "" {
		return false
	}

	f, err := os.Create(fileName)
	if err != nil {
		logger.Printf("pprof: could not open file for CPU profile: %s", err)
		return false
	}
	if err := pprof.StartCPUProfile(f); err != nil {
		logger.Printf("pprof: could not start CPU profile: %s", err)
		return false
	}
	logger.Printf("pprof: started CPU profile, save data to: %s", fileName)

	return true
}

func WriteMemProfile(fileName string, logger *util.Logger) {
	if fileName == "" {
		return
	}

	f, err := os.Create(fileName)
	if err != nil {
		logger.Printf("pprof: could not create memory profile: %s", err)
	}
	runtime.GC() // get up-to-date statistics
	if err := pprof.WriteHeapProfile(f); err != nil {
		logger.Printf("pprof: could not write memory profile: %s", err)
	}
	logger.Printf("pprof: wrote memory profile to %s", fileName)
	_ = f.Close()
}

