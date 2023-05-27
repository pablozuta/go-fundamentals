package main

import (
	"fmt"
	"math"
)

// declaramos struct con dos fields de tipo Int
type Point struct {
	X, Y int
}

// definimos un metodo que toma los valores del struct Point y devuelve un valor float64
func (p Point) Distance() float64 {
	// usamos un metodo del package math para sacar la raiz cuadrada de un valor (hacemos type-casting a float64 primero)
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

func main() {
	// creamos un objeto de typo Point
	p := Point{22, 33}
	// invocamos el metodo Distance usando como input el objeto p
	fmt.Println(p.Distance())
}
