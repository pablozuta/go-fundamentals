package main

import "fmt"

func main() {
	matrix := [][]int{
		{11, 12, 13},
		{22, 23, 24},
		{33, 34, 35},
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}
