package main

import "fmt"

func main()  {
	palabrasArray := [4] string {"manzana", "pera", "uva", "platano"}
	palabrasEnReversa := [4] string {}

	for i := len(palabrasArray) -1 ; i >= 0; i-- {
		palabrasEnReversa[len(palabrasArray)- i - 1] = palabrasArray[i]
	}  

	fmt.Printf("Array original %v\n", palabrasArray)
	fmt.Printf("Array en reversa %v\n", palabrasEnReversa)
}