package main

import (
	"fmt"
	"math"
)

func main() {
	const radius = 2.5

	circumference := 2 * math.Pi * radius
	area := math.Pi * math.Pow(radius, 2)

	fmt.Printf("The Circumference of the circle is %.2f\n", circumference)
	fmt.Printf("The Area of the circle is %.2f", area)
}
