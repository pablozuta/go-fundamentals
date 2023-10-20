package main

import (
	"fmt"
)

func performArithmeticOperations(x, y int) (int, int, int, float64) {
	sum := x + y
	difference := x - y
	product := x * y
	quotient := float64(x) / float64(y)
	return sum, difference, product, quotient
}

func main() {
	num1, num2 := 108, 17
	sum, difference, product, quotient := performArithmeticOperations(num1, num2)

	fmt.Printf("Numbers %d, %d\n", num1, num2)
	fmt.Printf("Sum %d\n", sum)
	fmt.Printf("Difference %d\n", difference)
	fmt.Printf("Product %d\n", product)
	fmt.Printf("Quotient %.2f\n", quotient)

}
