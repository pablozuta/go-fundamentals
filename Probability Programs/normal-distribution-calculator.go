package main

import (
	"fmt"
	"math"
)

func main() {
	mean := 0.0
	stdDev := 1.0

	x := 1.5
	probability := normalDistributionProbability(mean, stdDev, x)
	fmt.Printf("Probability of X <= %.2f in a standard normal distribution: %.4f\n", x, probability)

}
func normalDistributionProbability(mean, stdDev, x float64) float64 {
	z := (x - mean) / stdDev
	probability := 0.5 * (1 + math.Erf(z/math.Sqrt2))
	return probability
}
