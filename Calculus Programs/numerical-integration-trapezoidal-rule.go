// In this program, we define the integrateTrapezoidal function that approximates the definite integral of a function using the Trapezoidal Rule. The f function represents the integrand, and we pass the lower and upper limits of integration (a and b) as well as the number of subdivisions n as parameters to the integrateTrapezoidal function. The result is then printed.
package main

import (
	"fmt"
	"math"
)

// integrateTrapezoidal approximates the definite integral of a function using the Trapezoidal Rule.
func integrateTrapezoidal(f func(float64) float64, a float64, b float64, n int) float64 {
	h := (b - a) / float64(n)
	sum := (f(a) + f(b)) / 2.0

	for i := 1; i < n; i++ {
		x := a + float64(i)*h
		sum += f(x)
	}
	return sum * h
}

func main() {
	// define the function f(x) = x^2
	f := func(x float64) float64 {
		return math.Pow(x, 2)
	}

	// Calculate the definite integral of f(x) from 0 to 1 using the Trapezoidal Rule with 1000 subdivisions
	integral := integrateTrapezoidal(f, 0, 1, 1000)

	// mostrar el resultado
	fmt.Printf("The definitive integral of f(x) from 0 to 1 is approximately %.4f", integral)
}
