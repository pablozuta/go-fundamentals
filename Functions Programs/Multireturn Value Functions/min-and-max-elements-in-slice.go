package main

import "fmt"

// la funcion toma un slice de integers y devuelve el valor menor y el maximo
func minAndMax(numeros []int) (int, int) {
	// iniciamos las variables
	var min, max int
	// guardamos el primer elemento del array temporalmente como numero minimo
	min = numeros[0]
	for _, numero := range numeros {
		if numero > max {
			max = numero
		} else {
			min = numero
		}
	}
	return min, max
}

func main() {
	numeros := []int{1, 22, 33, 44, 55, 66, 2666}
	min, max := minAndMax(numeros)
	fmt.Println("Numero Minimo: ", min)
	fmt.Println("Numero Maximo: ", max)
}
