package main

import (
	"fmt"
)

type RectangularPrism struct {
	Length float64
	Width  float64
	Height float64
}

func (r RectangularPrism) Volume() float64 {
	return r.Length * r.Width * r.Height
}

func main() {
	prism := RectangularPrism{Length: 4.0, Width: 2.0, Height: 3.0}
	volume := prism.Volume()
	fmt.Printf("Rectangular Prism: Length = %.2f, Width = %.2f, Height = %.2f\n", prism.Length, prism.Width, prism.Height)
	fmt.Printf("Volume = %.2f\n", volume)

}
