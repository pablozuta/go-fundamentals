package main

import "fmt"

func main() {
	// define a 3x3 matrix
	matrix := [3][3]int {
		{11, 22, 3},
		{41, 56, 6},
		{7, 8, 98},
	}
	// print the matrix
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%4d", matrix[i][j])
		}
		fmt.Println()
	}

	// find the maximun element in the matrix
	max := matrix[0][0]
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] > max {
				max = matrix[i][j]
			}
		}
	}
	// print the maximun element
	fmt.Printf("The maximun element is %d", max)
}