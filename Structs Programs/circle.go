package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func(c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	circle :=Circle{radius: 5}
	fmt.Println("Area of the circle is:", circle.area())
}
