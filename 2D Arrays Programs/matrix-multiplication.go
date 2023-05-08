package main

import "fmt"

func main() {
	// define 2 matrix of 3x3
	matrixUno := [3][3]int{
		{2, 4, 6},
		{4, 7, 6},
		{9, 3, 5},
	}
	matrixDos := [3][3]int{
		{4, 6, 5},
		{1, 5, 7},
		{3, 6, 8},
	}

	// multiplicacion de las matrix
	var matrixMultipication [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrixMultipication[i][j] = matrixUno[i][j] * matrixDos[i][j]
		}
	}

	// mostramos los resultados
	for i := 0; i < len(matrixMultipication); i++ {
		for j := 0; j < len(matrixMultipication[i]); j++ {
			fmt.Print(matrixMultipication[i][j], " ")
		}
		fmt.Println()
	}

	// forma sencilla de mostrar el resultado
	fmt.Println(matrixMultipication)
}
