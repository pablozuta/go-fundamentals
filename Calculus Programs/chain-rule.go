package main

import (
	"fmt"
	"math"
)

// f1 represents the outer function in the composite function
func f1(x float64) float64 {
	return math.Sin(x)
}

// f2 represents the inner function in the composite function
func f2(x float64) float64 {
	return math.Pow(x, 2)
}

func main() {
	// Define the derivative of the outer function f1(x) = sin(x)
	derivativef1 := func(x float64) float64 {
		return math.Cos(x)
	}

	// Define the derivative of the inner function f2(x) = x^2
	derivativef2 := func(x float64) float64 {
		return 2 * x
	}

	// Calculate the derivative of the composite function using the chain rule
	x := 1.0
	derivative := derivativef1(f2(x)) * derivativef2(x)

	// Print the result
	fmt.Printf("The derivative of f(x) = sin(x^2) at x = 1 is approximately: %.4f\n", derivative)
}
