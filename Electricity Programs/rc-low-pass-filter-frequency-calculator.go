package main

import (
	"fmt"
	"math"
)

func freqCalculator(r, c float64) float64 {
	return 1.0 / (2 * math.Pi * r * c)
}

func main() {
	fmt.Println("RC Low-Pass Filter Cutoff Frequency Calculator")

	var r, c float64

	fmt.Print("Enter resistor value R1: ")
	fmt.Scanln(&r)
	fmt.Print("Enter capacitor value C: ")
	fmt.Scanln(&c)

	cutoffFrequency := freqCalculator(r, c)
	fmt.Printf("Cutoff Frequency: %.2f Hz\n", cutoffFrequency)

}
