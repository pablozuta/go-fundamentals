// The process can be executed asynchronously. This is done by calling the 'Start' method of the 'Cmd' structure. In this case, the process is executed, but the main goroutine does not wait until it ends. The 'Wait' method could be used to wait the process ends. After the 'Wait' method finishes, the resources of the process are released.

// This approach is more suitable for executing long-running processes and services that the program depends on.
package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	prc := exec.Command("pwd")
	// Create a buffer to store the output of the command.
	out := bytes.NewBuffer([]byte{})

	// Set prc's stdout to the created buffer.
	prc.Stdout = out
	err := prc.Start()
	if err != nil {
		fmt.Println(err)
	}

	prc.Wait()

	if prc.ProcessState.Success() {
		fmt.Println("Processs run successfully with output:")
		fmt.Println(out.String())
	}
}
