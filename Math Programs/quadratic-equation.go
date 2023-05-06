package main

import (
	"fmt"
	"math"
)

func main() {
	a := 1.0
	b := -5.0
	c := 6.0

	discriminant := math.Pow(b, 2) - 4 * a * c

	if discriminant < 0 {
		fmt.Println("The Equation Has No Real Solution")
	} else {
		x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		x2 := (-b - math.Sqrt(discriminant)) / (2 * a)

		fmt.Printf("The solutions to the quadratic equation %.2fx² + %.2fx + %.2f = 0 are %.2f and %.2f\n", a, b, c, x1, x2)
	}
}