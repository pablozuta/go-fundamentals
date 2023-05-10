package main

import "fmt"

//  recursively computes the Fibonacci numbers.
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// takes two arguments `from` and `to` that specify the range of Fibonacci numbers to compute. This function sends the computed Fibonacci numbers to a channel called `results`.
func computeFibonacci(from, to int, results chan int) {
	for i := from; i <= to; i++ {
		results <- fib(i)
	}
}

func main() {
	// we create a `results` channel with a capacity of 10
	results := make(chan int, 10)
	go computeFibonacci(0, 4, results)
	go computeFibonacci(5, 9, results)

	// we loop 10 times and print the computed Fibonacci numbers by receiving values from the `results` channel using the `<-` operator.
	for i := 0; i < 10; i++ {
		fmt.Println(<-results)
	}
}
