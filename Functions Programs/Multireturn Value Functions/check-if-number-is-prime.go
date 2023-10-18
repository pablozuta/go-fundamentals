package main

import (
	"fmt"
	"math"
)

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i < int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	testNumber := 17

	if isPrime(testNumber) {
		fmt.Printf("%d Is a Prime Number\n", testNumber)
	} else {
		fmt.Printf("%d Is NOT a Prime Number\n", testNumber)
	}
}
