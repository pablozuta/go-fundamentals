package main

import "fmt"

func main() {
	// EXPLICIT TYPE DECLARATION
	var numeroFloat float64 = 3.14159265359
	var numeroFloatDos float64 = 2.532513
	var sumaFloats float64 = numeroFloat + numeroFloatDos
	fmt.Println(sumaFloats)

	// NON EXPLICIT TYPE DECLARATIONS
	nFloat := 4.32
	nFloatDos := 9.642

	// ARITHMETIC OPERATIONS 
	suma := nFloat + float64(nFloatDos) // nFloatDos se convierte a float64
	resta := nFloat - float64(nFloatDos)
	multiplicacion := nFloat * float64(nFloatDos)
	division := nFloat / float64(nFloatDos)
	fmt.Println(suma)
	fmt.Println(resta)
	fmt.Println(multiplicacion)
	fmt.Println(division)



}