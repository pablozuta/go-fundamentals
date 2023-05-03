// An array in Go is a fixed-length data type that contains a contiguous
// block of elements of the same type.
// This could be a built-in type such as integers and strings
// or it can be a struct type.

package main

import "fmt"

func main() {
	// declaring arrays
	var arrayNumeros [5]int
	arrayNumeros[3] = 2666

	primes := [6]int {2, 3, 5, 7, 11, 13}

	var arrayFloats = [3] float32 {3.22, 311.4, 9.9444}

	// array de strings
	var arrayStrings [3]string
	arrayStrings[0] = "Dave"
	arrayStrings[1] = "Pamela"
	arrayStrings[2] = "George"

	// array de booleans
	var arrayBooleans [2]bool
	arrayBooleans[0] = true
	arrayBooleans[1] = false

	fmt.Println(arrayNumeros)
	fmt.Println(arrayFloats)
	fmt.Println(arrayStrings)
	fmt.Println(arrayBooleans)
	fmt.Println(primes)

	// declaring arrays using array literals
	arrayNumerosLiteral := [4]int {10, 30, 32, 643}
	fmt.Println(arrayNumerosLiteral)

	arrayStringsLiteral := [5]string {"Rethoric of David", "Oath of the Horatti", "The Parallel Lives", "Historie Ancienne", "Brutus Condemming his Sons to Death"}
	fmt.Println(arrayStringsLiteral[0])
	fmt.Println(arrayStringsLiteral[1])
	fmt.Println(arrayStringsLiteral[2])
	fmt.Println(arrayStringsLiteral[3])
	fmt.Println(arrayStringsLiteral[4])

	// array con longitud basada en cuantos elementos incluyamos
	var arrayNumerosLongitud = [...] int {2, 3, 7, 54, 32, 3}
	fmt.Println(arrayNumerosLongitud)

	// initialize specific index of an array
	// inicializa el index 1 con el valor 34 y el index 3 con el valor 60
	x := [5]int {1:34, 3:60} 
	fmt.Println("Inicializacion de Index 1 y 3")
	fmt.Println(x)

	// LOOPING EN ARRAYS
	for i:= 0; i < len(x); i++ {
		fmt.Println("looping por el array: ",x[i])
	}

	j := 0
	for range x {
		fmt.Println(x[j])
		j++
	}

	counter := 0
	for range arrayStrings {
		fmt.Println(arrayStrings[counter])
		counter ++
	}

	for i:= 0; i < len(arrayStrings); i++ {
		fmt.Println(arrayStrings[i])
	}
	// Capacity is determined based on the number of values initialized.
	// array con lenght definida de acuerdo a numero de elementos 
	vArray := [...]int {12, 3, 53, 87, 63, 525, 1}
	var vArray2 [7]int
	// copiamos el contenido de vArray a vArray2
	vArray2 = vArray
	fmt.Println(vArray)
	fmt.Println("vArray2: ", vArray2)

	// ver la longitud del array
	fmt.Println("Longitud del array vArray es: ",len(vArray))
	// asi accedemos al ultimo elemento del array
	fmt.Println("Accedemos al ultimo elemento de vArray: ", vArray[len(vArray) - 1])

}
// Arrays are valuable data structures because the memory is allocated sequentially.
// Having memory in a contiguous form can help to keep the memory you use 
// stay loaded within CPU caches longer.
// Using index arithmetic, you can iterate through all the elements
// of an array quickly.
// The type information for the array provides the distance in memory 
// you have to move to find each element.
// Since each element is of the same type and follows each other sequentially
// moving through the array is consistent and fast.