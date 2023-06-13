// This program simulates the birthday paradox, which states that in a group of people, the probability of at least two people having the same birthday is higher than one might intuitively expect. The program runs multiple simulations of a group of 23 people and checks if any two people share the same birthday. It then calculates the probability of this occurrence based on the simulation results.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// iniciamos una semilla random que usaremos para asignarla como dia de cumpleaños
	rand.Seed(time.Now().UnixNano())

	const numeroSimulaciones = 10000 // correremos el loop
	const numeroPersonas = 23
	var coincidenciasCumpleaños int

	for i := 0; i < numeroSimulaciones; i++ {
		// creamos un mapa 
		cumpleaños := make(map[int]bool)

		// generar un cumpleaño random para cada persona
		for j := 0; j < numeroPersonas; j++ {
			cumpleaño := rand.Intn(365) + 1
			if cumpleaños[cumpleaño] {
				coincidenciasCumpleaños++
				break
			}
			cumpleaños[cumpleaño] = true
		}
	}
	// calculamos la probabilidad de que un cumpleaño en el mismo dia suceda
	probabilidad := float64(coincidenciasCumpleaños) / float64(numeroSimulaciones)

	fmt.Printf("Probabilidad que al menos dos personas esten de cumpleaños el mismo dia en un grupo de %d es = %.2f\n", numeroPersonas, probabilidad)
	fmt.Printf("(Simulacion realizada %d veces)\n", numeroSimulaciones)

}
