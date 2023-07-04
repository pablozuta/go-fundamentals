package main

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

func distance(p1, p2 point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func main() {
	point1 := point{2.5, 3.5}
	point2 := point{1.0, 7.0}
	dist := distance(point1, point2)
	fmt.Printf("The distance between (%.2f, %.2f) and (%.2f, %.2f) is %.2f\n",
		point1.x, point1.y, point2.x, point2.y, dist)
}
