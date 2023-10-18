package main

import (
	"fmt"
	"math"
)

// funcion que retornara el area y circunferencia del circulo
func getCircleProps(radius float64) (float64, float64) {
	area := math.Pi * (math.Pow(radius, 2))
	circunference := 2 * (radius * math.Pi)
	return area, circunference
}

func main() {
	r := 5.0
	area, circunference := getCircleProps(r)

	fmt.Printf("The Area and Circunference of a circle with the radius of %.2f is :\n", r)
	fmt.Printf("Area = %.2f\n", area)
	fmt.Printf("Circunference = %.2f", circunference)

}
