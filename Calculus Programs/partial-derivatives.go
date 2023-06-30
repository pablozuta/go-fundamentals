package main

import (
	"fmt"
	"math"
)

// dfdx represents the partial derivative of f with respect to x
func dfdx(x, y float64) float64 {
	return 2 * x
}

// dfdy represents the partial derivative of f with respect to y
func dfdy(x, y float64) float64 {
	return math.Sin(y)
}

func main() {
	// Evaluate the partial derivatives at the point (1, π/4)
	x := 1.0
	y := math.Pi / 4

	// Calculate the partial derivatives
	partialDerivativeX := dfdx(x, y)
	partialDerivativeY := dfdy(x, y)

	// Print the results
	fmt.Printf("The partial derivative of f with respect to x at (1, π/4) is approximately: %.4f\n", partialDerivativeX)
	fmt.Printf("The partial derivative of f with respect to y at (1, π/4) is approximately: %.4f\n", partialDerivativeY)
}
