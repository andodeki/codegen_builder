package util


import (
	"crypto/ed25519"
	"os"
	"time"
	"os/exec"


	"github.com/andodeki/propertylisting/src/util"
)
func C_WFile(content, filename string) {
	f, createErr := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	// f, createErr := os.Create(filename)
	if createErr != nil {
		logger.Error("cannot create file:", createErr)
	}
	defer f.Close()

	_, writeErr := f.WriteString(time.Now().Format("2006-01-02 15:04:05") + "|" + content + "\n")
	if writeErr != nil {
		logger.Error("cannot write to file:", writeErr)
	}
}
func C_WFile2(content ed25519.PublicKey, filename string) {
	f, createErr := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	// f, createErr := os.Create(filename)
	if createErr != nil {
		logger.Error("cannot create file:", createErr)
	}
	defer f.Close()

	_, writeErr := f.Write(content)
	if writeErr != nil {
		logger.Error("cannot write to file:", writeErr)
	}
}
func ExecCmd(cmd string, args []string, logger *util.Logger) {
	output, err := executeCommand(cmd, args, logger)
	if err != nil {
		logger.Printf("bash shell command error: %v", err.Error())
	}
	logger.Infof("o: %v\n", output)
}

// https://pkg.go.dev/os/exec#Cmd.StdinPipe
func executeCommand(cmd string, args []string, logger *util.Logger) (string, error) {
	cmd_obj := exec.Command(cmd, args...)
	out, err := cmd_obj.Output()
	if err != nil {
		logger.Errorf("Output error: %v", err.Error())
	}
	return string(out), nil
}
func GreetingsMaker() string {

	hours, _, _ := time.Now().Clock()
	hourString := fmt.Sprintf("%d", hours)

	// fmt.Sprintf("%s", hours)
	// fmt.Println("hourString: ", hourString)
	hourInt, err := strconv.Atoi(hourString)
	if err != nil {
		return err.Error()
	}
	// fmt.Println("hourInt: ", hourInt)

	if hourInt >= 1 && hourInt &lt;= 12 {
		return "Good Morning"
	} else if hourInt >= 12 && hourInt &lt;= 16 {
		return "Good Afternoon"
	} else if hourInt >= 16 && hourInt &lt;= 21 {
		return "Good Evening"
	} else if hourInt >= 21 && hourInt &lt;= 24 {
		return "Good Evening"
	}

	return "Hello"

}

/*
=======================================
|| CONTEXT.go                        ||
=======================================
*/

/*
=======================================
|| ULID.go                        ||
=======================================
*/
/*
=======================================
|| ERRORS.go                        ||
=======================================
*/
/*
=======================================
|| LOGGER.go                        ||
=======================================
*/
