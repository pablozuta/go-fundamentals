// This program simulates a whistleblower's decision to reveal information and emphasizes the importance of ethical considerations, including protecting anonymity and safety.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// iniciamos una nueva semilla para el metodo rand
	rand.Seed(time.Now().UnixNano())

	// simulamos la decision de revelar informacion
	isWhistleBlower := rand.Float32() < 0.5 // Float32 genera un numero entre 0.0 y 1.0

	if isWhistleBlower {
		fmt.Println("You are a whistleblower.")
		fmt.Println("Ensure your anonimity and safety is protected.")
	} else {
		fmt.Println("You are NOT a whistleblower.")
		fmt.Println("Consider the ethical implications of your actions.")
	}
}
