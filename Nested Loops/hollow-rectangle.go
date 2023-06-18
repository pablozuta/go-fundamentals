package main

import "fmt"

func main() {
	rows := 5
	columns := 10
	for i := 1; i <= rows; i++ {
		for j := 1; j <= columns; j++ {
			if i == 1 || i == rows || j == 1 || j == columns {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
