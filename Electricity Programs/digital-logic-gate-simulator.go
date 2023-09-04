package main

import "fmt"

// funcion toma dos valores booleans y devuelve true solo si los dos valores son true (AND logic gate)
func AND(a, b bool) bool {
	return a && b
}

func main() {
	fmt.Println("Logical AND Gate Simulation")

	inputs := [][]bool{
		{false, false},
		{false, true},
		{true, false},
		{true, true},
	}

	fmt.Println("Inputs          | Outputs")
	fmt.Println("--------------------------")

	for _, input := range inputs {
		a, b := input[0], input[1]
		output := AND(a, b)
		fmt.Printf("%-6v | %v\n", input, output)
	}

}
