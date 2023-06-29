package main

import (
	"fmt"
)

// df represent the derivative of function f
func df2(x float64) float64 {
	return 2*x - 2
}

func gradientDescent2(initialX float64, learningRate float64, iterations int) float64 {
	x := initialX
	for i := 0; i < iterations; i++ {
		gradient := df2(x)
		x -= learningRate * gradient
	}
	return x
}

func main() {
	// set the initial guess , learning rate and number of iterations
	initialX := 0.0
	learningRate := 0.1
	iterations := 100

	// Perform gradient descent optimization
	optimalX := gradientDescent2(initialX, learningRate, iterations)

	// Print the result
	fmt.Printf("The optimal x that minimizes the function f(x) = x^2 - 2x + 1 is approximately: %.4f\n", optimalX)

}
