// This program simulates a basic analysis of exoplanet data. It generates data for a number of exoplanets, including their names, radii, and orbital periods. This data could be used for further analysis or visualization.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Exoplanet struct {
	Name      string
	Radius    float64
	OrbitDays float64
}

func main() {
	exoplanets := generateExoplanetData(10)

	fmt.Println("Exoplanet Data Program")
	fmt.Println("")
	for _, planet := range exoplanets {
		fmt.Printf("Planet: %s\n", planet.Name)
		fmt.Printf("Radius: %.2f\n", planet.Radius)
		fmt.Printf("Orbit Days %.2f days\n\n", planet.OrbitDays)
	}
}

func generateExoplanetData(count int) []Exoplanet {
	rand.Seed(time.Now().UnixNano())

	exoplanets := make([]Exoplanet, count)

	for i := 0; i < count; i++ {
		exoplanets[i] = Exoplanet{
			Name:      fmt.Sprintf("Exoplanet %d", i+1),
			Radius:    rand.Float64() * 10,
			OrbitDays: rand.Float64() * 1000,
		}
	}
	return exoplanets
}
