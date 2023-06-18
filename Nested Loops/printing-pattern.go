package main

import "fmt"

func main() {
	rows := 6
	for i := 1; i < rows; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("%2d", i)
		}
		fmt.Println()
	}
}
