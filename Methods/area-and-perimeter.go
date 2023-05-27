package main

import "fmt"

type Rectangulo struct {
	Largo float64
	Ancho float64
}

func(r Rectangulo) Area() float64 {
	return r.Largo * r.Ancho
}

func(r Rectangulo) Perimetro() float64 {
	return 2 * (r.Largo + r.Ancho) 
}

func main()  {
	rec := Rectangulo {22.4, 33.6}
	fmt.Println("Area:", rec.Area())
	fmt.Println("Perimetro:", rec.Perimetro())
}