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

func (t Triangle) Area1() float64 {
	semiperimeter := (t.Side1 + t.Side2 + t.Side3) / 2
	area := math.Sqrt(semiperimeter * (semiperimeter - t.Side1) * (semiperimeter - t.Side2) * (semiperimeter - t.Side3))
	return area

}
func (t Triangle) Perimeter() float64 {
	return t.Side1 + t.Side2 + t.Side3
}

func main() {
	triangle := Triangle{3.0, 4.0, 5.0}
	area := triangle.Area1()
	perimeter := triangle.Perimeter()
	fmt.Printf("Triangle: Side1 = %.2f, Side2 = %.2f, Side3 = %.2f\n", triangle.Side1, triangle.Side2, triangle.Side3)
	fmt.Printf("Area = %.2f\n", area)
	fmt.Printf("Perimeter = %.2f\n", perimeter)
}
