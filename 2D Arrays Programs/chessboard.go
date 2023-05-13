package main

import "fmt"

func main() {
	// define a 8x8 2D array
	board :=[8][8]string {
		{"R", "N", "B", "Q", "K", "B", "N", "R"},
		{"P", "P", "P", "P", "P", "P", "P", "P"},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{"p", "p", "p", "p", "p", "p", "p", "p"},
		{"r", "n", "b", "q", "k", "b", "n", "r"},
	}

	// print the chessboards
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print(board[i][j], "  ")
		}
		fmt.Println()
	}
}