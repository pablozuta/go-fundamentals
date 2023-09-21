// In this program, we implement optimization using the golden section search method. We define a function f(x) = x^2 - 2x + 1 that we want to minimize. The goldenSectionSearch function performs the optimization process by iteratively narrowing the search interval based on function evaluations. It continues until the interval becomes sufficiently small, determined by the epsilon value.  
  
// By providing the initial bounds (a and b) and the epsilon value, we can apply the golden section search method to find the optimal value of x that minimizes the function.
 
package main

import (
	"fmt"
	"math"
)

// f3 represents the function to optimize
func f3(x float64) float64 {
	return math.Pow(x, 2) - 2*x + 1
}

func goldenSectionSearch(a, b, epsilon float64) float64 {
	ratio := (3 - math.Sqrt(5)) / 2
	x1 := a + ratio*(b-a)
	x2 := b - ratio*(b-a)

	for math.Abs(b-a) > epsilon {
		if f3(x1) < f3(x2) {
			b = x2
			x2 = x1
			x1 = a + ratio*(b-a)
		} else {
			a = x1
			x1 = x2
			x2 = b - ratio*(b-a)
		}
	}
	return (a + b) / 2
}

func main() {
	// Set the initial bounds, epsilon (convergence criteria)
	a := 0.0
	b := 2.0
	epsilon := 0.0001

	// Perform optimization using the golden section search method
	optimalX := goldenSectionSearch(a, b, epsilon)

	// Print the result
	fmt.Printf("The optimal x that minimizes the function f(x) = x^2 - 2x + 1 is approximately: %.4f\n", optimalX)
}
