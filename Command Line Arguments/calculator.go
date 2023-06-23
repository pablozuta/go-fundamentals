package main

import (
	"fmt"
	"os"
	"strconv"
)
func main() {
	args := os.Args
	// parse the first number
	a, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		panic(err)
	}
	// parse the second number
	b, err := strconv.ParseFloat(args[3], 64)
	if err != nil {
		panic(err)
	}
	// perform the corresponding arithmetic operation based on the sign
	var result float64
	switch args[2] {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("Invalid Operator")
	}
	// print the result
	fmt.Println(result)
}