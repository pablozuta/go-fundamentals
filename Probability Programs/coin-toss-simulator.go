package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	const numTosses = 1000
	var heads, tails int

	for i := 0; i < numTosses; i++ {
		// generate a random float between 0 and 1 and check if it less than 0.5
		if rand.Float64() < 0.5 {
			// if random number is less than 0.5 increment the heads count
			heads++
		} else {
			// if random number is more than 0.5 increment the tails count
			tails++
		}
	}
	fmt.Printf("Of a total of %d Tosses\n", numTosses)
	// show occurrences of heads
	fmt.Printf("Heads: %d\n", heads)
	// show occurrences of tails
	fmt.Printf("Tails: %d\n", tails)
	// calculate and show the probability of getting heads
	fmt.Printf("Probability of Heads %.2f\n", float64(heads)/float64(numTosses))
	// calculate and show the probability of getting tails
	fmt.Printf("Probability of Tails %.2f\n", float64(tails)/float64(numTosses))
	
}
