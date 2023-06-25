// The Fibonacci sequence is a series of numbers in which each number is the sum of the two preceding numbers, starting from 0 and 1.
package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	n := 10
	fmt.Printf("The %d-th number in the fibonacci sequence is %d\n", n, fibonacci(n))
}

// In this example, the fibonacci function takes an integer n as input and returns the n-th number in the Fibonacci sequence. The function recursively calculates the n-th number by summing the (n-1)-th and (n-2)-th numbers in the sequence, until it reaches the base case of n=0 or n=1.
