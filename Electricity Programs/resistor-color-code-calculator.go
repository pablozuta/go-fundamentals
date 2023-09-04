package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var colorCodes = map[string]int{
	"black": 0, "brown": 1, "red": 2, "orange": 3,
	"yellow": 4, "green": 5, "blue": 6, "violet": 7,
	"gray": 8, "white": 9,
}

func main() {
	fmt.Println("Resistor Color Code Calculator")
	fmt.Println("Enter the colors of the bands separated by spaces (e.g., red black yellow gold):")

	reader := bufio.NewReader(os.Stdin)
	colors, _ := reader.ReadString('\n')

	colors = strings.TrimSpace(colors)
	bandColors := strings.Fields(colors)

	if len(bandColors) != 4 {
		fmt.Println("Invalid number of bands")
		return
	}
	significantDigits := colorCodes[bandColors[0]]*10 + colorCodes[bandColors[1]]
	multiplier := int(math.Pow(10, float64(colorCodes[bandColors[2]])))
	tolerance := bandColors[3]

	fmt.Printf("Resistance: %d ohms ± %s tolerance\n", significantDigits*multiplier, tolerance)

}
