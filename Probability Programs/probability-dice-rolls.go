package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	const numRolls = 10000
	const targetCombination = 7

	var count int
	for i := 0; i < numRolls; i++ {
		roll1 := rand.Intn(6) + 1
		roll2 := rand.Intn(6) + 1

		if roll1+roll2 == targetCombination {
			count++
		}
	}
	probability := float64(count) / float64(numRolls)
	fmt.Printf("Probability of getting a sum  of %d with two dice rolls %.4f\n", targetCombination, probability)
}
