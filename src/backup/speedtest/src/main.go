package main



import (
    "github.com/andodeki/propertylisting/src/runsvc"
)
func main() {
	cmdOptions, cmdName := runsvc.GetCmdOptions()
	var conf = runsvc.GetConfig(cmdOptions, cmdName)
	Application(conf, cmdOptions)
}
