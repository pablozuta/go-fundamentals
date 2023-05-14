package main

import "fmt"

func main() {
	// 2D array de integers
	matrixInt := [3][3]int{
		{2, 4, 6},
		{3, 6, 9},
		{1, 8, 1},
	}

	// mostrar el primer nodo
	fmt.Println("primer nodo", matrixInt[0][0])
	// mostrar el ultimo nodo
	fmt.Println("ultimo nodo", matrixInt[2][2])

	// mostrar sin formato
	fmt.Println(matrixInt)

	// la longitud de la matrix
	fmt.Println("Longitud de matrix:", len(matrixInt)) // output 3

	// nested for loop para mostrar todos los nodos
	for i := 0; i < len(matrixInt); i++ {
		for j := 0; j < len(matrixInt[i]); j++ {
			fmt.Print(matrixInt[i][j], "  ")
		}
		fmt.Println()
	}

	fmt.Println("Valores en reversa")
	// nested loop para mostrar los valores en reversa
	for i := len(matrixInt) - 1; i >= 0; i-- {
		for j := len(matrixInt[i]) - 1; j >= 0; j-- {
			fmt.Print(matrixInt[i][j], "  ")
		}
		fmt.Println()
	}

}
