package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("Capacitor Charging Simulator")
	fmt.Println("Press Control+C to Exit")

	voltageSource := 5.0
	resistor := 1000.0
	capacitance := 0.0001
	timeConstant := resistor * capacitance

	charge := 0.0

	for {
		charge += (voltageSource - charge) * (1 - math.Exp(-1.0/timeConstant))
		fmt.Printf("Time: %.2f seconds, Charge: %.4f C\n", time.Since(time.Time{}).Seconds(), charge)
		time.Sleep(100 * time.Millisecond)

	}
}
