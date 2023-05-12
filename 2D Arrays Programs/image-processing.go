package main

import "fmt"

func main() {
	// define a 4x4 grayscale
	image := [4][4] int {
		{10, 20, 30, 40},
		{50, 60, 70, 80},
		{90, 100, 110, 120},
		{130, 140, 150, 160},
	}

	// invert the colors of the image
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[i]); j++ {
			image[i][j] = 255 - image[i][j]
		}
	}

	// print the inverted image
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[i]); j++ {
			fmt.Print(image[i][j], " ")
		}
		fmt.Println()
	}
}