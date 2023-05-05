// This program takes an array of strings as input and concatenates all the strings in the array into a single string, separated by commas.  
package main

import "fmt"


func main() {
	// creamos un array de strings
	arr := [] string {"Apple", "Peach", "Watermelon", "Berry"}
	// inicializamos una variable de tipo string para guardar los resultados
	var concatenated string

	// creamos un for loop que ira de principio a final del array
	for i := 0; i < len(arr); i++ {
		// si el index del array es 0 guardamos ese valor en la variable "concatenated"
		if i == 0 {
			concatenated = arr[i]
		} else {
			// para el resto de los indices concatenamos el valor de la variable con el indice del array
			concatenated = concatenated + ", " + arr[i] 
		}
	}
	fmt.Printf("Concatenated string: %s\n", concatenated)
}