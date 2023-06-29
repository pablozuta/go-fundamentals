package main

import (
	"fmt"
	"math"
)

func main() {
	const angleDegrees = 45.0

	angleRadians := angleDegrees * (math.Pi / 180.0)
	sine := math.Sin(angleRadians)

	fmt.Printf("The sine of %.2f degrees is %.4f", angleDegrees, sine)
}
