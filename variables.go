package main

import "fmt"

func main() {
	// declaring variables without type using the 'var' keyword

	var numero = 100 // integer (there are many types of integer types in GO)
	var nombre = "Pedro Paramo" // para strings se usan double quotes
	var george = true // boolean
	var floating = 100.5 // float (there are float32 and float64)
	var myRune = 'A' // this is a rune (Unicode character, returns a UTF-( number))
	fmt.Println(numero)
	fmt.Println(nombre)
	fmt.Println(george)
	fmt.Println(floating)
	fmt.Println(myRune)
	fmt.Println("")

	// there are many types of integers in GO divided into Signed / Unsigned

	// SIGNED (4 categories)
	var numeroDos int8 = -109 // range -128 to 127
	var numeroDos2 int8 = 120

	var numeroTres int16 = 5637 // range -32768 to 32767
	var numeroTres2 int16 = -32400

	var numeroCuatro int32 = -281738464 // range -2147483648 to 2147483647
	var numeroCuatro2 int32 = 5462334

	var numeroCinco int64 = 283748491375 // range -9223372036854775808 to 9223372036854775807
    var numeroCinco2 int64 = -4325854343

	fmt.Println("----SIGNED-----")
	fmt.Println("este es un int8: ", numeroDos)
	fmt.Println("este es un int8: ", numeroDos2)
	fmt.Println("este es un int16: ", numeroTres)
	fmt.Println("este es un int16: ", numeroTres2)
	fmt.Println("este es un int32: ", numeroCuatro)        
	fmt.Println("este es un int32: ", numeroCuatro2)        
	fmt.Println("este es un int64: ", numeroCinco)
	fmt.Println("este es un int64: ", numeroCinco2)
	fmt.Println("")

	// UNSIGNED (4 categories) solo numeros positivos (y 0)

	var numeroSeis uint8 = 233 // range  0 to 255
	var numeroSeis2 uint8 = 176

	var numeroSiete uint16 = 4555 // range 0 to 65535
	var numeroSiete2 uint16 = 8763

	var numeroOcho uint32 = 159384732 // range 0 to 4294967295
	var numeroOcho2 uint32 = 2313421

	var numeroNueve uint64 = 3574637281634 // range 0 to 18446744073709551615
    var numeroNueve2 uint64 = 746272816374

	fmt.Println("----UNSIGNED-----")
	fmt.Println("este es un uint8: ", numeroSeis)
	fmt.Println("este es un uint8: ", numeroSeis2)
	fmt.Println("este es un int16: ", numeroSiete)
	fmt.Println("este es un int16: ", numeroSiete2)
	fmt.Println("este es un int32: ", numeroOcho)
	fmt.Println("este es un int32: ", numeroOcho2)
	fmt.Println("este es un int64: ", numeroNueve)
	fmt.Println("este es un int64: ", numeroNueve2)

	// FLOATS are float32 y float64 (el default es float64)
	var miFloatDefault = 3.543 // sera de tipo float64
	var miFloat32 float32 = 4.321
	var miFloat64 float64 = 543.33
	fmt.Println(miFloatDefault)
	fmt.Println(miFloat32)
	fmt.Println(miFloat64)

	// ver el tipo de dato de una variable
	fmt.Printf("The value of miFloatDefault es %T\n", miFloatDefault)
	fmt.Printf("The value of miFloat32 es %T\n", miFloat32)

	// booleans
	var miBooleanVerdad bool = true
	var miBooleanMentira bool = false
	fmt.Println(miBooleanVerdad)
	fmt.Println(miBooleanMentira)

	// runes
	var miRuna rune ='♥'
	fmt.Println(miRuna) // output 9829
	fmt.Printf("%c\n", miRuna) // output ♥

	// complex numbers
	var unComplex64 complex64 = 3 + 4i
	fmt.Println(unComplex64) // output: (3 + 4i)
	var unComplex128 complex128 = complex(7, 4)
	fmt.Println(unComplex128) // output (7 + 4i)

	// strings
	var unString = "American Farm"
	fmt.Println(unString)

	// DECLARACIONES CORTAS
	cUno := 20
	cDos := 43.2
	cTres := "LOOPS"
	cCuatro := true
	fmt.Println(cUno)
	fmt.Println(cDos)
	fmt.Println(cTres)
	fmt.Println(cCuatro)

	// CONSTANTES
	const unaConstante string = "Tales Of Mystery And Imagination"
	const unaConstanteInt int = 23

	// constantes con implicit data type
	const constanteNumero = 300
	const constanteBool = false
	const constanteString = "it was the shape of the GALLOWS!"
	const constanteFloat = 9.8371
	fmt.Println(constanteString)


	
}