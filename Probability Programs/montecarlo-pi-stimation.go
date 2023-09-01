package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const totalPoints = 1000000
	var insideCircle int

	for i := 0; i < totalPoints; i++ {
		x := rand.Float64()
		y := rand.Float64()

		distance := x*x + y*y
		if distance < 1 {
			insideCircle++
		}
	}
	estimatedPI := 4 * float64(insideCircle) / float64(totalPoints)
	fmt.Printf("Estimated value of PI using monte carlo simulation: %.6f\n", estimatedPI)
}
