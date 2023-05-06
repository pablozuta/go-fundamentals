package main

import (
	"fmt"
	"math"
)

func main() {
	x := 5.0
	y := -10.0

	fmt.Println("Ablosute value of x:", math.Abs(x))
	fmt.Println("Power of x raised to 2:", math.Pow(x, 2))
	fmt.Println("Square root of x:", math.Sqrt(x))
	fmt.Println("Sine of y:", math.Sin(y))
	fmt.Println("Cosine of y:", math.Cos(y))
	fmt.Println("Tangent of y:", math.Tan(y))
}
