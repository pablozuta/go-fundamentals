// This program creates a 2D array of integers and initializes its elements with multiplication table values.  
package main

import "fmt"

func main() {
	// creamos un array 2D(matrix) con 10 columnas y 10 filas
	multiplicationTable := [10][10] int {}
	
	// inicializamos los valores de cada nodo de la matrix
	// usamos la variable i para el indice de columnas y la variable j para el indice de las filas
	for i := 0; i < 10; i ++ {
		for j := 0; j < 10; j++ {
			multiplicationTable[i][j] = (i+1) * (j + 1)
		}
	}
	
	// print the multiplication table
	for i := 0; i < 10; i ++ {
		for j := 0; j < 10; j++ {
			// el %4 es una metrica para la separacion de las columnas
			fmt.Printf("%4d", multiplicationTable[i][j])
			
		}
		fmt.Println() // lo usamos para crear un salto de linea despues de 10 columnas
	}


}