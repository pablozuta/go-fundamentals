package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	// iniciamos una semilla para el generador de numeros random
	rand.Seed(time.Now().UnixNano())
	// constante para usar en el for loop principal
	const numRolls = 1000
	// iniciamos un array de 6 elementos a su zero-value
	var frequencies [6]int
	
	
}