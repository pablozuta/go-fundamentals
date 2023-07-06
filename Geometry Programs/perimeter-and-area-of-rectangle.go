package main

import "fmt"

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func main() {
	rect := Rectangle{Length: 5.0, Width: 3.0}
	perimeter := rect.Perimeter()
	area := rect.Area()
	fmt.Printf("Rectangle: Length = %.2f, Width = %.2f\n", rect.Length, rect.Width)
	fmt.Printf("Perimeter = %.2f\n", perimeter)
	fmt.Printf("Area = %.2f\n", area)

}
