package main

import (
	"fmt"
	"math"
)

func main() {
	radius := 5.0
	area := math.Pi * math.Pow(radius, 2)
	fmt.Printf("The area of the circle with radius %.2f is %.2f", radius, area)
}
