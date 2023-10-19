// the Go binary could also be used as a tool for various utilities and with use of the go run as a replacement for the bash script.
// for these purposes, it is usual that the command-line utilities are called.
package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	// Command represents an external command being prepared or run.
	prc := exec.Command("ls", "-a") // Create a new command to execute "ls -a".

	// Create a buffer to store the output of the command.
	out := bytes.NewBuffer([]byte{})

	// Set prc's stdout to the created buffer.
	prc.Stdout = out

	// Run starts the specified command and waits for it to complete.
	err := prc.Run()

	// Check if there was an error during command execution.
	if err != nil {
		fmt.Println(err)
	}

	// Check if the process executed successfully.
	if prc.ProcessState.Success() {
		fmt.Println("Process run successfully with output:")
		fmt.Println(out.String()) // Print the output of the executed command.
	}
}
