package main

import "fmt"

func main() {
	const numero int16 = 30
	const PI = 3.14159
	const nombre string = "Pablo"

	// Numeric constants have no size or sign
	// can be of arbitrary high precision and do no overflow
	const numeroLargo = 0.838273858584285

	// multiple assignments
	const Lunes, Martes, Miercoles = 1, 2, 3

	// enums
	const (
		Hombre = 1
		Mujer = 2
	)

	// explicit type declarations
	const flotador float32 = 9.876
	const verdadero bool = true
	const complejo complex128 = 1 + 2i

	fmt.Println(numero)
	fmt.Println(PI)
	fmt.Println(nombre)
	fmt.Println(numeroLargo)
	fmt.Println("dias:", Lunes, Martes, Miercoles)
	fmt.Println(flotador)
	fmt.Println(verdadero)
	fmt.Println(complejo)
}