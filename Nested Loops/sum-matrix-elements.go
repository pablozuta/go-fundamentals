package main

import "fmt"

func main() {
	matrix := [][]int{
		{2, 5, 6},
		{4, 5, 8},
		{2, 1, 6},
	}
	rows := len(matrix)
	columns := len(matrix[0])
	suma := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			suma += matrix[i][j]
		}
	}
	fmt.Println("Resultado suma matrix es:", suma)
}
