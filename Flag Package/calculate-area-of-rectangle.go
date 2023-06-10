package main

import (
	"fmt"
	"flag"
)

func main()  {
	// definir command line flags
	widthPtr := flag.Float64("width", 0.0, "width of rectangle")
	heightPtr := flag.Float64("height", 0.0, "height of rectangle")

	// parsing
	flag.Parse()

	// calcular el area
	area := *widthPtr * *heightPtr
	fmt.Printf("El area del rectangulo es: %.2f ", area)
	

	// podemos correr el programa asi:
	// go run calculate-area-of-rectangle.go -width=10 -height=5  
}