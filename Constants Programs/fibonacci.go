package main

import "fmt"

func main() {
	const limit = 10
	fmt.Println("Fibonacci Sequence:")
	fmt.Print("0 ")

	a, b := 0, 1
	for i := 2; i < limit; i++ {
		next := a + b
		fmt.Printf("%d ", next)
		a = b
		b = next
	}
}
