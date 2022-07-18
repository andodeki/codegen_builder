//AUTO GENERATED By gen_main.go, DO NOT EDIT
package main

import (
	"context"
	"os"
	"runtime/pprof"

	"github.com/andodeki/api.gen_test.com/src/config"
	"github.com/andodeki/api.gen_test.com/src/datasources"
	"github.com/andodeki/api.gen_test.com/src/queue"
	"github.com/andodeki/api.gen_test.com/src/util"
)

// is set through linker by build.sh
var buildVersion string
var buildTime string

func Application(cfg *config.Config) {
	cmdOptions, cmdName := getCmdOptions()
	logger := util.NewLogger(*cfg)

	if cfg.LogWorkerStart {
		logger.Infof("main: start test_1 version=%s", buildVersion)
	}

	// start cpu profiling if enabled
	if runCpuProfile(string(cmdOptions.CpuProfile)) {
		defer pprof.StopCPUProfile()
	}

	exitCode := func() (exitCode int) {
		// whenever an error is pushed to this chan, main is terminated
		initiateShutdown := make(chan error, 4)

		var client *datasources.DatabaseClient
		var scyllaClient *datasources.ScyllaDBClient
		var redisClient *datasources.RedisClient

		emailInAppQueue := queue.NewEmailQueue(logger, cfg)

		promExporter, errPromE := datasources.NewOTExporter(context.Background(), logger, cfg, initiateShutdown)
		if errPromE != nil {
			logger.Infof("main: datasources.NewOTExporter: %!s(MISSING)", errPromE)
		}

		httpServerInstance := runHttpServer(logger, cfg, emailInAppQueue, client, redisClient, scyllaClient, promExporter)
		defer httpServerInstance.Shutdown()

		if cfg.LogWorkerStart {
			logger.Print("main: start completed; run until SIGTERM or SIGINT is received")
		}
	}()
	if cfg.LogWorkerStart {
		logger.Infof("main: stutdown completed; exit %!d(MISSING)", exitCode)
	}
	os.Exit(exitCode)
}