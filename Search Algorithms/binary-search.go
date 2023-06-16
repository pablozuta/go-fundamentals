// Binary search is a more efficient search algorithm that works by repeatedly dividing the search interval in half. It requires the input array to be sorted beforehand.
package main

import (
	"fmt"
)

// la funcion toma como argumento un array de integers y un numero integer y devuelve tambien como resultado  un integer
func binarySearch2(arr []int, target int) int {
	fmt.Println("Cantidad de Elementos del Array:", len(arr))
	low := 0
	high := len(arr) - 1

	// comenzamos un for loop hasta que el indice low y high sean iguales(lo que quiere decir que no hay mas elementos en los cuales buscar)
	// indice low se mantiene en 0 siempre
	for low <= high {
		mid := (low + high) / 2
		fmt.Printf("Index Low %d ", low)
		fmt.Printf("Index High %d ", high)
		fmt.Printf("Index Mid %d ", mid)

		// empezamos buscando en el elemento medio del array
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			fmt.Println("Mid Index es MENOR a =", target)
			low = mid + 1
		} else {
			fmt.Println("Mid Index es MAYOR a Target=", target)
			high = mid - 1
		}
	}
	return -1
}

func main() {
	// el array debe estar ordenado
	arr := []int{1, 5, 7, 9, 12, 15, 44}

	target := 7
	result := binarySearch2(arr, target)

	if result == -1 {
		fmt.Println("Elemento no encontrado")
	} else {
		fmt.Printf("Elemento encontrado en index %d", result)
	}

}
