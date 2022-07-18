package main

import (
	"context"
	"os"
	"runtime/pprof"
	"os/signal"
	"syscall"
	"time"

	"github.com/andodeki/propertylisting/src/runsvc"
	"github.com/andodeki/propertylisting/src/datasources"
	"github.com/andodeki/propertylisting/src/queue"
	"github.com/andodeki/propertylisting/src/util"
    "github.com/andodeki/propertylisting/src/config"
)
func Application(cfg *config.Config, cmdOptions runsvc.CmdOptions) {
	//cmdOptions, cmdName := runsvc.GetCmdOptions()
	logger := util.NewLogger(*cfg)

	if cfg.LogWorkerStart {
		logger.Infof("main: start test_1 version=%s", runsvc.BuildVersion)
	}

	// start cpu profiling if enabled
	if runsvc.RunCpuProfile(string(cmdOptions.CpuProfile)) {
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
        // setup SIGTERM, SIGINT handlers
		gracefulStop := make(chan os.Signal, 1)
		signal.Notify(gracefulStop, syscall.SIGTERM)
		signal.Notify(gracefulStop, syscall.SIGINT)

		// wait for something to trigger a shutdown
		select {
		case err := <-initiateShutdown:
			logger.Infof("main: forced shutdown due to fatal error: %s", err)
			exitCode = runsvc.ExitDueToModuleStart
		case sig := <-gracefulStop:
			if cfg.LogWorkerStart {
				logger.Infof("main: graceful shutdown; caught signal: %+v", sig)
			}
			exitCode = runsvc.ExitSuccess
		}

		// for emailInAppQueue.Size() > 0 {
		// 	time.Sleep(time.Millisecond * 500)
		// }

		for emailInAppQueue.Size() > 0 {
			time.Sleep(time.Millisecond * 500)
		}
		// write memory profile; after that defer will run the shutdown methods
		runsvc.WriteMemProfile(string(cmdOptions.MemProfile), logger)

		return

	}()
	if cfg.LogWorkerStart {
		logger.Infof("main: stutdown completed; exit %!d(MISSING)", exitCode)
	}
	os.Exit(exitCode)
}
