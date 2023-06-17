package main

import (
	"fmt"
	"math"
)

// calculates the numerical approximation of the derivative of a function at a given point.
func derivativeApproximation(f func(float64) float64, x float64, h float64) float64 {
	return (f(x+h) - f(x)) / h
}

func main() {
	// define the function
	f := func(x float64) float64 {
		return math.Pow(x, 2)
	}
	// calculate the derivative approximation
	derivative := derivativeApproximation(f, 2, 0.001)
	// show the result
	fmt.Printf("The derivative of f(x) at x = 2 is approximatly %.4f", derivative)
}
