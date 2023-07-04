package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// generate 10 random numbers from a normal distribution
	for i := 0; i < 10; i++ {
		num := generateNormalDistribution(0, 1)
		fmt.Printf("Random number from normal distribution %.2f\n", num)
	}
}

func generateNormalDistribution(mean, stdDev float64) float64 {
	// Use the Box-Muller transform to generate two independent standard normal random variables
	u1 := rand.Float64()
	u2 := rand.Float64()
	z1 := math.Sqrt(-2*math.Log(u1)) * math.Cos(2*math.Pi*u2)

	num := mean + stdDev*z1
	return num
}
