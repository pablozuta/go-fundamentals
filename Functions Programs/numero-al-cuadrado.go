package main

import "fmt"

func numeroAlCuadrado(n int)int  {
	return n * n
}

func main() {
	n := 5
	fmt.Printf("%d al cuadrado es : %d", n, numeroAlCuadrado(n))
}