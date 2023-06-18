package main

import "fmt"

func main() {
	rows := 6

	for i := 0; i < rows; i++ {
		for j := i; j >= 1; j-- {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
}
