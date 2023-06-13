// This program simulates the Monty Hall problem, a probability puzzle. It runs multiple simulations where the player chooses a door, the host reveals a door without the prize, and then the player decides whether to stay with their initial choice or switch to the other unopened door. The program calculates the probabilities of winning by staying or switching based on the simulation results.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	const numSimulations = 1000
	var stayWins, switchWins int

	for i := 0; i < numSimulations; i++ {
		// randomly assign the winning door
		winningDoor := rand.Intn(3) + 1
		// player initial choice
		initialChoice := rand.Intn(3) + 1

		// host reveals a door that does not have the prize and was not chosen by the player
		var revealedDoor int
		for {
			revealedDoor := rand.Intn(3) + 1
			if revealedDoor != winningDoor && revealedDoor != initialChoice {
				break
			}
		}
		// player final choice after switching doors
		finalChoice := 6 - initialChoice - revealedDoor
		// check if player wins by staying
		if initialChoice == winningDoor {
			stayWins++
		}
		// check if player wins by switching
		if finalChoice == winningDoor {
			switchWins++
		}

	}
	stayWinProbability := float64(stayWins) / float64(numSimulations)
	switchWinProbability := float64(switchWins) / float64(numSimulations)

	fmt.Printf("Probability of winning by staying: %.2f\n", stayWinProbability)
	fmt.Printf("Probability of winning by switching: %.2f\n", switchWinProbability)
}
