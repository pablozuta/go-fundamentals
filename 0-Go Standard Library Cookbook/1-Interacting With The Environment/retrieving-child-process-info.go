// This recipe shows how to obtain the PID and elementary information about the child process after it terminates.
package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	// creamos una variable para usar dependiendo del Sistema Operativo
	var cmd string

	if runtime.GOOS == "windows" {
		cmd = "timeout"
	} else {
		cmd = "sleep"
	}
	proc := exec.Command(cmd, "1")
	proc.Start()
	// no process state is returned till the process finish.
	fmt.Printf("Process state for running process: %v\n", proc.ProcessState)
	// the PID could be obtain even for the running process
	fmt.Printf("PID of running process: %d\n\n", proc.Process.Pid)
	// cada vez que se ejecuta el programa el PID sera diferente
}
