 
//AUTO GENERATED By gen_main.go, DO NOT EDIT
package main



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