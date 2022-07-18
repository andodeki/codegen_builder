package gen_conf

import (
	"fmt"

	"github.com/andodeki/gen_project/aOK_GenerateCode/dir"
	"github.com/andodeki/gen_project/aOK_GenerateCode/project"
)

const (
	header = ` 
//AUTO GENERATED By gen_main.go, DO NOT EDIT
package main

import (
	%s
)



func getConfig(cmdOptions CmdOptions, cmdName string) *config.Config {
	// read, transform and validate configuration
	cfg, err := config.ReadConfigFile(cmdName, string(cmdOptions.Config))
	if len(err) > 0 {
		for _, e := range err {
			log.Printf("config: error: %v", e)
		}
		os.Exit(ExitDueToConfig)
	}

	if cfg.LogConfig {
		if err := cfg.PrintConfig(); err != nil {
			log.Printf("config: cannot print: %s", err)
		}
	}

	return &cfg
}

func getCmdOptions() (cmdOptions CmdOptions, cmdName string) {
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
		fmt.Println("github.com/andodeki/api.propertylist.com version:", buildVersion)
		fmt.Println("build at:", buildTime)
		os.Exit(ExitSuccess)
	}

	return cmdOptions, parser.Name
}
`
)

var ConfImports = []string{
	"src/config",
}

func Generate_ConfigContents() {
	sb := dir.FormStringDir(project.ProjectName, ConfImports)

	configuration_ := fmt.Sprintf(header, sb.String())
	dir.Create_Files(configuration_, "configuration.go")

}