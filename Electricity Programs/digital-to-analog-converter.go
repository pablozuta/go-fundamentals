package main

import (
	"fmt"
	"math"
)

func dacOutputVoltage(binary int64, bits int, referenceVoltage float64) float64 {
	return (float64(binary) / math.Pow(2, float64(bits))) * referenceVoltage
}

func main() {
	fmt.Println("Digital-to-Analog Converter (DAC) Output Voltage Calculator")

	var binary int64
	var bits int
	var referenceVoltage float64

	fmt.Print("Enter Binary Input: ")
	fmt.Scanln(&binary)
	fmt.Print("Enter Number of Bits: ")
	fmt.Scanln(&bits)
	fmt.Print("Enter Reference Voltage: ")
	fmt.Scanln(&referenceVoltage)

	outputVoltage := dacOutputVoltage(binary, bits, referenceVoltage)
	fmt.Printf("Analog Output Voltage: %.2f V\n", outputVoltage)

}
