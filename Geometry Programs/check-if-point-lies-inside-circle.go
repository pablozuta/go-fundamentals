package main

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

func isPointInsideCircle(p point, center point, radius float64) bool {
	distance := math.Sqrt(math.Pow(p.x-center.x, 2) + math.Pow(p.y-center.y, 2))
	return distance <= radius
}

func main() {
	circleCenter := point{x: 0, y: 0}
	circleRadius := 5.0
	testPoint := point{x: 3, y: 4}

	inside := isPointInsideCircle(testPoint, circleCenter, circleRadius)
	fmt.Printf("Circle: Center(%v, %v), Radius=%.2f\n", circleCenter.x, circleCenter.y, circleRadius)
	fmt.Printf("Test Point: (%v, %v)\n", testPoint.x, testPoint.y)

	if inside {
		fmt.Println("The point lies inside the circle.")
	} else {
		fmt.Println("The point lies outside the circle.")
	}
}
