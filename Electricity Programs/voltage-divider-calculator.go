// This program calculates the output voltage of a voltage divider circuit based on the resistor values and input voltage.
package main

import "fmt"

func voltageDivider(inputVoltage, r1, r2 float64) float64 {
	return inputVoltage * (r1 / (r1 + r2))
}

func main() {
	fmt.Println("Voltage Divider Calculator")
	var inputVoltage, r1, r2 float64

	fmt.Print("Enter input voltage (V): ")
	fmt.Scanln(&inputVoltage)
	fmt.Print("Enter resistance R1 (Ω): ")
	fmt.Scanln(&r1)
	fmt.Print("Enter resistance R2 (Ω): ")
	fmt.Scanln(&r2)

	outputVoltage := voltageDivider(inputVoltage, r1, r2)
	fmt.Printf("Output Voltage: %.2f V\n", outputVoltage)

}
