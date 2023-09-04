package main

import (
	"fmt"
)

func main() {
	var option int
	fmt.Println("Ohm's law calculator")
	fmt.Println("1. Calculate Voltage")
	fmt.Println("2. Calculate Current")
	fmt.Println("3. Calculate Resistance")
	fmt.Println("Enter your choice (1/2/3): ")
	fmt.Scanln(&option)

	switch option {
	case 1:
		var current, resistance float64
		fmt.Print("Enter current (A): ")
		fmt.Scanln(&current)
		fmt.Print("Enter resistance (Ω): ")
		fmt.Scanln(&resistance)
		voltage := current * resistance
		fmt.Printf("Voltage: %.2f V\n", voltage)
	case 2:
		var voltage, resistance float64
		fmt.Print("Enter voltage (V): ")
		fmt.Scanln(&voltage)
		fmt.Print("Enter resistance (Ω): ")
		fmt.Scanln(&resistance)
		current := voltage / resistance
		fmt.Printf("Current: %.2f A\n", current)
	case 3:
		var voltage, current float64
		fmt.Print("Enter voltage (V): ")
		fmt.Scanln(&voltage)
		fmt.Print("Enter current (A): ")
		fmt.Scanln(&current)
		resistance := voltage / current
		fmt.Printf("Resistance: %.2f Ω\n", resistance)

	default:
		fmt.Println("Invalid Choice")
	}
}
