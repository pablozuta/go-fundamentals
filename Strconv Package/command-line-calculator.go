package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage calculator: <number> <operator> <number>")
		os.Exit(1)
	}

	a, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Invalid First Argument", os.Args[1])
		os.Exit(1)
	}

	b, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Invalid Third Argument", os.Args[3])
		os.Exit(1)
	}
	var result float64
	switch os.Args[2] {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "/":
		result = a / b
	case "*":
		result = a * b
	default:
		fmt.Println("Invalid Operator Sign")
	}
	fmt.Println(result)

}
