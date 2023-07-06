package main

import (
	"fmt"
	"math"
)

type Triangle struct {
	Side1 float64
	Side2 float64
	Side3 float64
}

func (t Triangle) Area() float64 {
	semiperimeter := (t.Side1 + t.Side2 + t.Side3) / 2
	area := math.Sqrt(semiperimeter * (semiperimeter - t.Side1) * (semiperimeter - t.Side2) * (semiperimeter - t.Side3))
	return area
}

func main() {
	triangle := Triangle{Side1: 5.0, Side2: 7.0, Side3: 9.0}
	area := triangle.Area()
	fmt.Printf("Triangle: Side1=%.2f, Side2=%.2f, Side3=%.2f\n", triangle.Side1, triangle.Side2, triangle.Side3)
	fmt.Printf("Area = %.2f\n", area)
}
