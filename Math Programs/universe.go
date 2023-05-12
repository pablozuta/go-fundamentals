package main

import (
	"fmt"
	"math"
)

func main() {
	// define constants of distance in meters
	const (
		lightYear = 9.461e15
		parsec = 3.086e16
		au = 1.496e11
	)
	// calcular la distancia al sistema de estrellas mas cercano (Proxima Centaury) en metros
	nearestStarDist := 4.24 * lightYear
	fmt.Printf("Distance to Proxima Centauri %.1f meters\n", nearestStarDist)

	// calcular el diametro de la Via Lactea en metros
	milkyWayDiameter := 100000 * parsec
	fmt.Printf("Diameter of the Milky Way galaxy: %.1f meters\n", milkyWayDiameter)

	// calcular la edad del universo en segundos
	ageOfUniverse := 13.8 * math.Pow(10, 9) * 365 * 24 * 60 * 60
	fmt.Printf("Age of the Universe in seconds: %.0f\n", ageOfUniverse)
}
