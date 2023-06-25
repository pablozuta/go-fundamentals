// The factorial of a non-negative integer n is the product of all positive integers less than or equal to n.
package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	n := 5
	fmt.Printf("The factorial of %d is %d\n", n, factorial(n))
}

// In this example, the factorial function takes an integer n as input and returns the factorial of n. The function recursively calculates the factorial by multiplying n with the factorial of (n-1), until it reaches the base case of n=0.
