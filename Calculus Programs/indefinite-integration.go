// In this program, we define the function f(x) = sin(x) that we want to integrate indefinitely. The indefiniteIntegration function performs indefinite integration using the given lower and upper limits of integration (a and b) and the number of subdivisions n. It approximates the integral using the Riemann sum method.

// By providing the appropriate inputs, we can compute the indefinite integral of the function.

package main

import (
	"fmt"
	"math"
)

// the function to integrate
func f4(x float64) float64 {
	return math.Sin(x)
}

func indefiniteIntegration(a, b float64, n int) float64 {
	deltaX := (b - a) / float64(n)
	sum := 0.0

	for i := 0; i < n; i++ {
		x := a + float64(i)*deltaX
		sum += f4(x) * deltaX
	}
	return sum
}

func main() {
	// Set the integration limits and the number of subdivisions
	a := 0.0
	b := math.Pi
	n := 1000

	// Perform indefinite integration using the given limits and subdivisions
	result := indefiniteIntegration(a, b, n)

	// Print the result
	fmt.Printf("The indefinite integral of sin(x) is approximately: %.4f\n", result)
}
