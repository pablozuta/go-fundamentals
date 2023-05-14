// Linear search, also known as sequential search, is a simple search algorithm that searches for an element in a list or an array in a linear manner. It checks each element of the list one by one until a match is found.
package main

import "fmt"

// la funcion toma dos argumentos ; un array de integers y un integer
func linearSeach(arr []int, target int) int {
	// un for loop comparara cada elemento del array con el 'target value'
	for i := 0; i < len(arr); i++ {
		// si los elementos son iguales devuelve el indice del elemento del array
		if arr[i] == target {
			return i
		}
	}
	// si termina el for loop sin hacer match devolvemos un valor -1 que usaremos para mostrar un mensaje que elemento no fue encontrado
	return -1
}

func main() {
	// declaramos e inicializamos un array de integer
	arr := []int{1, 4, 6, 7, 34, 3}
	// inicializamos el target value
	target := 34
	// almacenamos el return value de la funcion en una variable
	result := linearSeach(arr, target)

	// usamos el valor de retorno de la funcion para mostrar mensajes
	if result == -1 {
		fmt.Println("Elemento no encontrado")
	} else {
		fmt.Printf("Elemento encontrado en indice %d", result)
	}

}
