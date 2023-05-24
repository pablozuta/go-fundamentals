package main

import "fmt"

func changeValue(ptr *int, newValue int) {
	*ptr = newValue
}

func main()  {
	var x int
	var ptr *int

	ptr = &x
	*ptr = 10
	fmt.Printf("Direccion del puntero = %p\n", ptr)
	fmt.Printf("Valor del puntero (derreferenciacion) = %d\n", *ptr)
	changeValue(ptr, 20)
	fmt.Printf("Valor de x despues de usar la funcion changeValue = %d", x)
}