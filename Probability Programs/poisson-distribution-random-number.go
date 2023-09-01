package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const (
		lambda     = 2.5
		numSamples = 10_000
	)
	counts := make(map[int]int)

	for i := 0; i < numSamples; i++ {
		poissonValue := generatePoisson(lambda)
		counts[poissonValue]++
	}
	fmt.Println("Poisson Distribution (λ = 2.5) Simulation:")
	for value, frequency := range counts {
		probability := float64(frequency) / float64(numSamples)
		fmt.Printf("Value: %d, Frequency: %d, Probability: %.4f\n", value, frequency, probability)
	}
}

func generatePoisson(lambda float64) int {
	p := 1.0
	k := 0

	for p > math.Exp(-lambda) {
		k++
		p *= rand.Float64()
	}
	return k - 1
}
