package main

import "fmt"

func main() {
	// creamos un byte slice
	bs :=[]byte {'H', 'o', 'l', 'a'}
	// print the slice
	fmt.Println(bs) // 71 111 108 97
	// modificar un caracter del slice
	bs[0] = 'S'
	// mostrar en forma de string
	fmt.Println(string(bs))
	// hacer un append de un elemento
	bs = append(bs, '!') 
	fmt.Println(string(bs))
}