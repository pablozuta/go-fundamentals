package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// iniciamos una semilla para el generador de numeros random
	rand.Seed(time.Now().UnixNano())
	// constante para usar en el for loop principal
	const numRolls = 1000
	// iniciamos un array de 6 elementos a su zero-value
	var frequencies [6]int

	// generamos un numero random entre 0 y 5 y agregamos 1 para obtener un numero entre 0 y 6 (inclusive)
	for i := 0; i < numRolls; i++ {
		roll := rand.Intn(6) + 1
		// incrementamos la cuenta en el elemento correspondiente del array
		frequencies[roll-1]++
	}

	for i, frequency := range frequencies {
		// calculamos la probabilidad al dividir el numero por el numero de rolls
		probability := float64(frequency) / float64(numRolls)
		// mostramos el numero, la frecuencia y la probabilidad
		fmt.Printf("Numero %d: Frecuencia: %d, Probabilidad: %.2f\n", i+1, frequency, probability)
	}

}
