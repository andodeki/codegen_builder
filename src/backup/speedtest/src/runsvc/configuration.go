package runsvc

//Configuration//
import (
	"fmt"
	"log"
	"os"
    

	"github.com/andodeki/propertylisting/src/config"
	"github.com/jessevdk/go-flags"

)
// is set through linker by build.sh
var BuildVersion string
var BuildTime string

const (
	ExitSuccess          = 0
	ExitDueToCmdOptions  = 1
	ExitDueToConfig      = 2
	ExitDueToModuleStart = 3
)
type CmdOptions struct {
	Version    bool           `long:"version" description:"Print the build version and timestamp"`
	Config     flags.Filename `short:"c" long:"config" description:"Config File in yaml format" default:"./config.yaml"`
	CpuProfile flags.Filename `long:"cpuprofile" description:"write cpu profile to <file>"`
	MemProfile flags.Filename `long:"memprofile" description:"write memory profile to <file>"`
}

func GetConfig(cmdOptions CmdOptions, cmdName string) *config.Config {
	// read, transform and validate configuration
	cfg, err := config.ReadConfigFile(cmdName, string(cmdOptions.Config))
	if len(err) > 0 {
		for _, e := range err {
			log.Printf("config: error: %!v(MISSING)", e)
		}
		os.Exit(ExitDueToConfig)
	}

	if cfg.LogConfig {
		if err := cfg.PrintConfig(); err != nil {
			log.Printf("config: cannot print: %!s(MISSING)", err)
		}
	}

	return &cfg
}

func GetCmdOptions() (cmdOptions CmdOptions, cmdName string) {
	// parse command line options
	parser := flags.NewParser(&cmdOptions, flags.Default)
	parser.Usage = "[-c <path to yaml config file>]"
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(ExitSuccess)
		} else {
			os.Exit(ExitDueToCmdOptions)
		}
	}

	if cmdOptions.Version {
		fmt.Println("github.com/andodeki/propertylisting version:", BuildVersion)
		fmt.Println("build at:", BuildTime)
		os.Exit(ExitSuccess)
	}

	return cmdOptions, parser.Name
}
