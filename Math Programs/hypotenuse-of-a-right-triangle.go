package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3.0
	b := 4.0
	c := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	fmt.Printf("The Hypotenuse of a right triangle with sides %.2f and %.2f is %.f2\n", a, b, c)
}