package main

import (
	"fmt"
	"math"
)

type Point3D struct {
	X float64
	Y float64
	Z float64
}

func distance3D(p1, p2 Point3D) float64 {
	return math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2) + math.Pow(p2.Z-p1.Z, 2))
}

func main() {
	point1 := Point3D{X: 1.0, Y: 2.0, Z: 3.0}
	point2 := Point3D{X: 4.0, Y: 5.0, Z: 6.0}
	dist := distance3D(point1, point2)
	fmt.Printf("Point 1: (%.2f, %.2f, %.2f)\n", point1.X, point1.Y, point1.Z)
	fmt.Printf("Point 2: (%.2f, %.2f, %.2f)\n", point2.X, point2.Y, point2.Z)
	fmt.Printf("Distance = %.2f\n", dist)

}
