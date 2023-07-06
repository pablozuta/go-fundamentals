package main

import (
	"fmt"
	"math"
)

type Sphere struct {
	Radius float64
}

func (s Sphere) Volume() float64 {
	return (4.0 / 3.0) * math.Pi * math.Pow(s.Radius, 3)
}

func main() {
	sphere := Sphere{Radius: 5.0}
	volume := sphere.Volume()
	fmt.Printf("Sphere: Radius = %.2f\n", sphere.Radius)
	fmt.Printf("Volume = %.2f\n", volume)
}
