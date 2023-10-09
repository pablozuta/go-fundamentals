package main

import "fmt"

func bubbleSort(collection []int) {
	n := len(collection) - 1

	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			if collection[j] > collection[j+1] {
				collection[j], collection[j+1] = collection[j+1], collection[j]
			}
		}
	}
}

func main() {
	numeros := []int{66, 108, 55, 44, 33, 22, 11, 9}
	// mostramos el orden original de nuestra coleccion
	fmt.Println(numeros)

	// invocamos la funcion
	bubbleSort(numeros)
	fmt.Println(numeros)

}
