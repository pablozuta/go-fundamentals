package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const numSamples = 1000000

	average := calculateAverage(numSamples)
	fmt.Printf("Average of %d random numbers: %.4f\n", numSamples, average)

}

func calculateAverage(numSamples int) float64 {
	sum := 0.0
	for i := 0; i < numSamples; i++ {
		sum += rand.Float64()
	}
	average := sum / float64(numSamples)
	return average
}
