// Signals are the elementary way the operating systems communicate with the running process. Two of the most usual signals are called SIGINT and SIGTERM. These cause the program to terminate.
// There are also signals such as SIGHUP.SIGHUP indicates that the terminal which called the process was closed and, for example, the program could decide to move to the background.
// Go provides a way of handling the behavior in case the application received the signal.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// create the channel where the received signal would be sent.
	// the Notify will not block when the signal is sent and the channel s not ready.
	// so it is better to create a buffered channel
	sChan := make(chan os.Signal, 1)

	// Notify will catch the given signals and send the os.Signal value through the sChan.
	// if no signal is specified in the argument, all signals are matched.
	signal.Notify(sChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	// create a channel to wait till the signal is handled.
	exitChan := make(chan int)
	go func() {
		signal := <-sChan
		switch signal {
		case syscall.SIGHUP:
			fmt.Println("The calling terminal has been closed")
			exitChan <- 0
		case syscall.SIGINT:
			fmt.Println("The process has been interrupted by CTRL+C")
			exitChan <- 1
		case syscall.SIGTERM:
			fmt.Println("Kill SIGTERM was executed for process")
			exitChan <- 1
		case syscall.SIGQUIT:
			fmt.Println("Kill SIGQUIT was executed for process")
			exitChan <- 1
		}
	}()

	code := <-exitChan
	os.Exit(code)
}

// run the program
// send the SIGINT signal by pressing CTRL+C (secondary CTRL key on windows)
