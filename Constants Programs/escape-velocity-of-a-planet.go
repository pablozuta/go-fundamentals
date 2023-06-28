package main

import (
	"fmt"
	"math"
)

func main() {
	const (
		GravitationalConstant = 6.67430e-11 // m^3 kg^−1 s^−2
		PlanetMass            = 5.97e24     // kg (earth mass)
		PlanetRadius          = 6.371e6     // m (earth radius)
	)
	escapeVelocity := math.Sqrt((2 * GravitationalConstant * PlanetMass) / PlanetRadius)
	fmt.Printf("The escape velocity of the planet is %.2f m/s", escapeVelocity)
}
