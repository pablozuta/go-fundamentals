package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const (
		startingCapital = 100
		targetCapital   = 200
		winProbability  = 0.5
	)
	var numSimulations = 10000
	var bankruptCount int

	for i := 0; i < numSimulations; i++ {
		capital := startingCapital

		for capital > 0 && capital < targetCapital {
			if rand.Float64() < winProbability {
				capital++
			} else {
				capital--
			}
		}
		if capital == 0 {
			bankruptCount++
		}
	}
	probabilityOfBankruptcy := float64(bankruptCount) / float64(numSimulations)
	fmt.Printf("Probability of going bankrupt starting with $%d and aiming for $%d: %.4f\n", startingCapital, targetCapital, probabilityOfBankruptcy)
}
