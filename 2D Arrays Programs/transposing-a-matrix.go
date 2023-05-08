package main

import "fmt"

func main() {
	// define a 3x3 array
	matrix := [3][3]int {
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// print the original matrix
	fmt.Println("Original Matrix")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
	
	// transpose the matrix
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix[i]); j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}
	
	// print the transposed matrix
	fmt.Println("Transposed Matrix")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}

}