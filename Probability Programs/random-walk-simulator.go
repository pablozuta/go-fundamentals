package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// iniciamos la semilla random
	rand.Seed(time.Now().UnixNano())
	// numero de veces que se repetira el loop
	const numeroPasos = 100
	// variable de tipo int para guardar los resultados
	var posicion int

	for i := 1; i < numeroPasos; i++ {
		// genera un numero 1 o -1 (no genera 0)
		direccionPaso := rand.Intn(2)*2 - 1
		// se usa el numero generado para sumarlo al valor de la variable posicion
		posicion += direccionPaso

		fmt.Printf("Paso %d Posicion %d\n", i, posicion)
	}

}
