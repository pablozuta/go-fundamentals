// This program simulates a scenario where an individual needs to make an ethical decision and encourages them to consider ethical principles when making their choice.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// iniciamos una nueva semilla para el metodo rand
	rand.Seed(time.Now().UnixNano())

	// Simulate an scenario where an ethical decision needs to be made
	ethicalChoice := rand.Float32() < 0.5 // Float32 genera un numero entre 0.0 y 1.0

	if ethicalChoice {
		fmt.Println("You have chosen the ethical option.")
		fmt.Println("Considerations of ethical principles are essential.")
	} else {
		fmt.Println("You have chosen an unethical option.")
		fmt.Println("Reflect on the consequences of your decision.")
	}
}
