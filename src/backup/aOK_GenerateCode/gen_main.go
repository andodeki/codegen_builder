package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/andodeki/gen_project/aOK_GenerateCode/dir"
	"github.com/andodeki/gen_project/aOK_GenerateCode/gen_conf"
	"github.com/andodeki/gen_project/aOK_GenerateCode/gen_script"
	"github.com/andodeki/gen_project/aOK_GenerateCode/gen_startup"
	"github.com/andodeki/gen_project/aOK_GenerateCode/gen_utils"
)

const (
	header = ` 
//AUTO GENERATED By gen_main.go, DO NOT EDIT
package main


%s

const (
	ExitSuccess          = 0
	ExitDueToCmdOptions  = 1
	ExitDueToConfig      = 2
	ExitDueToModuleStart = 3
)

func main() {
	cmdOptions, cmdName := getCmdOptions()
	var conf = getConfig(cmdOptions, cmdName)
	Application(conf)
}
`
)

var MainImports = []string{
	"src/config",
	"src/startup",
}

func main() {
	// sb := FormStringDir(project.ProjectName, dir.MainImports)

	dir.Generate_Dir()

	Generate_MainContents()
	gen_utils.Generate_LoggerContents()
	gen_conf.Generate_ConfigContents()
	gen_startup.Generate_StartupContents()
	gen_script.Generate_ScriptContents()
	// bashShell()
	// var conf = Configuration.Get_configuration()
	// Startup.Application.Build(conf)
}

func Generate_MainContents() {
	s := CmdOptions()
	main_ := fmt.Sprintf(header, s)

	dir.Create_Files(main_, "main.go")
}

func bashShell() {
	// cmdExec := fmt.Sprintf("go mod init && go mod tidy")

	// command := []string{
	// 	"/<path>/yourscript.sh",
	// 	"arg1=val1",
	// 	"arg2=val2",
	// }

	// _, _ = Execute("/<path>/yourscript.sh", command)

	// os.Chdir(fmt.Sprintf("../%s", project.ProjectName))
	_, err := exec.Command("/bin/sh", "../api.gen_test.com/scripts/gomod.sh").Output()
	// cmd.Dir = fmt.Sprintf("../%s", project.ProjectName)

	// err := cmd.Run()
	if err != nil {
		log.Printf("shell command error:%v", err.Error())
	}
}

func Execute(script string, command []string) (bool, error) {

	cmd := &exec.Cmd{
		Path:   script,
		Args:   command,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	log.Printf("Executing command %s", cmd)

	err := cmd.Start()
	if err != nil {
		return false, err
	}

	err = cmd.Wait()
	if err != nil {
		return false, err
	}

	return true, nil
}

func CmdOptions() (s string) {
	// s =  fmt.Sprint(`type CmdOptions struct {
	// 	Version bool `long:"version" description:"Print the build version and timestamp"`
	// 	Config     flags.Filename `short:"c" long:"config" description:"Config File in yaml format" default:"./config.yaml"`
	// 	CpuProfile flags.Filename `long:"cpuprofile" description:"write cpu profile to <file>"`
	// 	MemProfile flags.Filename `long:"memprofile" description:"write memory profile to <file>"`
	// }`)

	return s
}

func generateStruct(name, typef, format string, fieldNames []string) {
	st := `type %s struct{
		%s %s %s
	}`

	for _, f := range fieldNames {

	}
}
