package main

import (
	"fmt"
	"math"
)

type Sphere1 struct {
	Radius float64
}

func (s Sphere1) SurfaceArea() float64 {
	return 4 * math.Pi * math.Pow(s.Radius, 2)
}

func (s Sphere1) Volume() float64 {
	return (4.0 / 3.0) * math.Pi * math.Pow(s.Radius, 3)
}

func main() {
	sphere := Sphere1{3.0}
	surfaceArea := sphere.SurfaceArea()
	volume := sphere.Volume()
	fmt.Printf("Sphere: Radius = %.2f\n", sphere.Radius)
	fmt.Printf("Surface Area = %.2f\n", surfaceArea)
	fmt.Printf("Volume = %.2f\n", volume)
}
