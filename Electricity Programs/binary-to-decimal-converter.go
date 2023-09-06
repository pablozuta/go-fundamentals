// This program converts a binary number to its decimal equivalent, which is useful in digital electronics.

package main

import (
	"fmt"
	"math"
)

func binaryToDecimal(binary int64) int64 {
	decimal := int64(0)
	position := 0

	for binary != 0 {
		remainder := binary % 10
		decimal += remainder * int64(math.Pow(2, float64(position)))
		binary /= 10
		position++
	}
	return decimal

}

func main() {
	fmt.Println("Binary to Decimal Converter")

	var binary int64

	fmt.Print("Enter binary number: ")
	fmt.Scanln(&binary)

	decimal := binaryToDecimal(binary)
	fmt.Printf("Decimal Equivalent: %d\n", decimal)
}
