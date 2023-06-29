package main

import "fmt"

func main() {
	const (
		length = 5.0
		width  = 3.0
	)

	area := length * width
	fmt.Printf("The area of the rectangle is %.2f", area)
}
