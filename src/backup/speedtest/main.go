package main

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
	"unsafe"
)

type Record struct {
	FirstName       string `db:"first_name"`
	LastName        string `db:"last_name"`
	Address         string `db:"address"`
	PictureLocation string `db:"picture_location"`
}

func main() {
	record := Record{}
	_ = unsafe.Sizeof(record)

	bashShell()

}

func bashShell() {
	// cmd := "cat"
	// args := []string{"/etc/os-release"}
	// osVersion(cmd, args)
	// hostcmd("ssh", []string{"-t", "-p2900", "root@127.0.0.1", "hostname"})
	execcmd("pwd", []string{})
	// aio := execcmd("ssh", []string{"-t", "-p2900", "root@127.0.0.1", "cat", "/proc/sys/fs/aio-max-nr"})
	// aioint, _ := strconv.Atoi(strings.Replace(aio, "\n", "", -1))

	// x := 1000048576
	// // fmt.Printf("%t\n%q\n", aioint == x, strings.Replace(aio, "\n", "", -1))
	// if aioint != x {
	// 	execcmd("ssh", []string{"-t", "-p2900", "root@127.0.0.1", "sudo", "sysctl", "-w", "fs.aio-max-nr=1000048576"})
	// } else if aioint == x {
	// 	execcmd("ssh", []string{
	// 		"ubuntu@10.38.195.126", "-A", "sudo", "systemctl", "start", "scylla-server",
	// 	})
	// 	execcmd("ssh", []string{
	// 		"ubuntu@10.38.195.20", "-A", "sudo", "systemctl", "start", "scylla-server",
	// 	})
	// 	execcmd("ssh", []string{
	// 		"ubuntu@10.38.195.158", "-A", "sudo", "systemctl", "start", "scylla-server",
	// 	})
	// }
	// // lxc exec db1-dev -- scp
	// // sudo sysctl -w fs.aio-max-nr=1000048576
	// // 1000048576

}
func execcmd(cmd string, args []string) string {
	output, err := executeCommand(cmd, args)
	if err != nil {
		log.Printf("bash shell command error: %v", err.Error())
	}
	fmt.Printf("o: %v\n", output)
	return output
}
func lxcls(cmd string, args []string) {
	output, err := executeCommand(cmd, args)
	if err != nil {
		log.Printf("bash shell command error: %v", err.Error())
	}
	fmt.Printf("o: %v\n", output)
}
func osVersion(cmd string, args []string) {
	output, err := executeCommand(cmd, args)
	if err != nil {
		log.Printf("bash shell command error: %v", err.Error())
	}
	fmt.Printf("o: %v\n", output)
}

func executeCommand(cmd string, args []string) (string, error) {
	// https://pkg.go.dev/os/exec#Cmd.StdinPipe
	// subProcess := exec.Cmd{
	// 	Args: args,
	// }
	cmd_obj := exec.Command(cmd, args...)
	// cmd_obj.Stdout = os.Stdout

	out, err := cmd_obj.Output()
	if err != nil {
		log.Printf("CombinedOutput error: %v", err.Error())
	}

	return string(out), nil
}

// func main() {
// 	// slice := []string{"a", "b", "c", "d", "e"}
// 	start := time.Now()
// 	sliceLength := 1000000
// 	var wg sync.WaitGroup
// 	wg.Add(sliceLength)
// 	fmt.Println("Running for loop…")
// 	for i := 0; i < sliceLength; i++ {
// 		go func(i int) {
// 			defer wg.Done()
// 			// val := slice[i]
// 			fmt.Printf("i: %v\n", i)
// 		}(i)
// 	}
// 	fmt.Println("Doing other stuff")
// 	wg.Wait()
// 	timeElapsed := time.Since(start)
// 	fmt.Printf("took [%s]", timeElapsed)
// 	fmt.Println("Finished for loop")
// }

// var wg sync.WaitGroup
// 	go loopOver(&wg)
// 	// wg.Done()
// 	timeElapsed := time.Since(start)
// 	fmt.Printf("took [%s]", timeElapsed)
// 	wg.Wait()

func loopOver(wg *sync.WaitGroup) {
	for i := 0; i < 1000000; i++ {
		defer wg.Done()
		wg.Add(1)
		fmt.Println("go")

	}

}
