// This program is a "Fair Random Number Generator" that generates a random number between 1 and 10. It emphasizes fairness and the absence of biases when generating random numbers, which is an ethical consideration in various applications, such as games and contests.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// iniciamos una nueva semilla para el metodo rand
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Welcome to the fair number generator")
	fmt.Println("Generating a random number between 1 and 10...")

	// guardamos el resultado de la funcion en una variable
	randomNumber := generateFairRandomNumber(1, 10)
	// mostramos el resultado por pantalla
	fmt.Printf("Random Number: %d\n", randomNumber)
}

// funcion que generara el numero random entre 1 y 10
func generateFairRandomNumber(min, max int) int {
	if min >= max {
		return min
	}
	return rand.Intn(max-min+1) + min
}
