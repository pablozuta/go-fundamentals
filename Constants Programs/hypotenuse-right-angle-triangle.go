package main

import (
	"fmt"
	"math"
)

func main() {
	const (
		sideA = 3.0
		sideB = 4.0
	)

	hypotenuse := math.Sqrt(math.Pow(sideA, 2) + math.Pow(sideB, 2))
	fmt.Printf("The Hypotenuse of the right angle triangle is %.2f", hypotenuse)
}
