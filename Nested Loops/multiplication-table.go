package main

import "fmt"

func main() {

	for i := 1; i < 11; i++ {
		for j := 1; j < 11; j++ {
			fmt.Printf("%3d ", i*j)
		}
		fmt.Println()
	}
}
