// (funcionara al ejecutarse en GOPATH)
// Getting to know the PID of the running process is useful. The PID could be used by OS utilities to find out the information about the process itself.
// It is also valuable to know the PID in case of process failure, so you can trace the process behavior across the system in system logs, such as /var/log/messages, /var/log/syslog.

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	// The function Getpid from the os package returns the PID of a process.
	pid := os.Getpid()
	fmt.Printf("Process PID: %d \n", pid)

	prc := exec.Command("ps", "-p", strconv.Itoa(pid), "-v")
	out, err := prc.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
