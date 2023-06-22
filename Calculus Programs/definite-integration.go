// In this program, we define the function f(x) = sin(x) that we want to integrate. The definiteIntegration function performs definite integration using the given lower and upper limits of integration (a and b) and the number of subdivisions n. It approximates the integral using the Riemann sum method.
// By providing the appropriate inputs, we can compute the definite integral of the function over the specified interval.
package main

import (
	"fmt"
	"math"
)

// f represent the function to integrate
func f(x float64) float64 {
	return math.Sin(x)
}
func definiteIntegration(a float64, b float64, n int) float64 {
	deltaX := (b - a) / float64(n)
	sum := 0.0

	for i := 0; i < n; i++ {
		x := a + float64(i)*deltaX
		sum += f(x) * deltaX
	}
	return sum
}

func main() {
	// set the lower and upper limits of integration and the number of subdivisions
	a := 0.0
	b := math.Pi
	n := 1000
	// perform definite integration using the given limits and subdivisions
	integral := definiteIntegration(a, b, n)

	// print the results
	fmt.Printf("The definite integral of sin(x) from 0 to π is approximately: %.4f\n", integral)
}
