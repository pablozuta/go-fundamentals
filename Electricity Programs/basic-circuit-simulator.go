package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Basic Circuit Simulator")
	fmt.Println("Press Control+C to Stop")

	voltageSource := 5.0
	resistance := 1000.0

	for {
		current := voltageSource / resistance

		fmt.Printf("Voltage Source: %.2f V, Resistance: %.2f Ω\n", voltageSource, resistance)
		fmt.Printf("Current: %.4f A\n", current)
		fmt.Println()

		// wait for a short duration
		time.Sleep(1000 * time.Millisecond)
	}
}
