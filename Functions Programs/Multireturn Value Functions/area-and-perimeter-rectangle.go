package main

import "fmt"

func calculateRectangleProps(length, width float64) (float64, float64) {
	area := length * width
	perimeter := 2 * (length + width)
	return area, perimeter
}

func main() {
	l, w := 5.0, 3.0
	area, perimeter := calculateRectangleProps(l, w)
	fmt.Printf("The area of the rectangle of length %.2f and width %.2f is :\n", l, w)
	fmt.Printf("Area = %.2f\n", area)
	fmt.Printf("Perimeter = %.2f\n", perimeter)

}
