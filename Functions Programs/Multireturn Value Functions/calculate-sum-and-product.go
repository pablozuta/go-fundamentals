package main

import (
	"fmt"
)

func calculateSumAndProduct(a, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}

func main() {
	num1, num2 := 20, 3
	sum, product := calculateSumAndProduct(num1, num2)
	fmt.Printf("Sum of numbers %d and %d is %d\n", num1, num2, sum)
	fmt.Printf("Product of numbers %d and %d is %d\n", num1, num2, product)

}
