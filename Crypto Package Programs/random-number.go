package main

import (
	"crypto/rand"
	"fmt"
	
)

func main() {
	// esto generara un array con la cantidad de numeros que especifiquemos dentro de []
	var buf [10]byte
	_, err := rand.Read(buf[:])
    if err != nil {
		panic(err)
	}
    // mostramos la variable buf como array
	fmt.Println(buf)

	// mostramos la variable buf numero por numero
	for i := 0; i < len(buf); i++ {
		fmt.Printf("%v\n", buf[i])
	}
}