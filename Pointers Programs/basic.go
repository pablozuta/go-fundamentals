package main

import "fmt"

func main() {
	var x int
	var pointer *int

	pointer = &x
	*pointer = 10

	fmt.Printf("Valor de x: %d\n", x)
	fmt.Printf("Valor de pointer: %p", pointer) // output 0xc0000a2058

}
