// In this program, we define a function f(x) = x^2 + 2x + 1 that we want to minimize using gradient descent. The derivative of f is calculated in the df function. The gradientDescent function performs the optimization process by iteratively updating the value of x based on the gradient and learning rate.
// By providing an initial guess for x, a learning rate, and the number of iterations, we can apply gradient descent to find the optimal value of x that minimizes the function.

package main

import (
	"fmt"
	"math"
)

// f represent the function to optimize
func f(x float64) float64 {
	return math.Pow(x, 2) + 2*x + 1
}

// df represents the derivative of the function f
func df(x float64) float64 {
	return 2*x + 2
}

func gradientDescent(initialX float64, learningRate float64, iterations int) float64 {
	x := initialX

	for i := 0; i < iterations; i++ {
		gradient := df(x)
		x -= learningRate * gradient

	}
	return x
}
func main() {
	// set the initial guess, learning rate and number of iterations
	initialX := 0.0
	learningRate := 0.1
	iterations := 100

	// perform gradient descent optimization
	optimalX := gradientDescent(initialX, learningRate, iterations)

	// print the results
	fmt.Printf("The optimal x that minimizes the function f(x) = x^2 + 2x + 1 is approximately: %.4f\n", optimalX)
}
