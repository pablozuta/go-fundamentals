package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("LED Blinking Simulator")
	fmt.Println("Enter Ctrl+C to Exit")

	ledState := false

	for {
		if ledState {
			fmt.Println("LED OFF")
			ledState = false
		} else {
			fmt.Println("LED ON")
			ledState = true
		}
		time.Sleep(1000 * time.Millisecond)
	}
}
