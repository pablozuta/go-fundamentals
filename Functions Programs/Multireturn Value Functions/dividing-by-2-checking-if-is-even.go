package main

import "fmt"

// la funcion toma un argumento de tipo int y devuelve un valor int
// y un valor booleano que representa si el argumento es par
func divideByTwoCheckIfEven(num int) (int, bool) {
	return num / 2, num%2 == 0
}

func main() {
	num := 7
	// guardamos el resultado de la funcion en dos variables
	result, isEven := divideByTwoCheckIfEven(num)

	if isEven {
		fmt.Printf("%d is Even, half of it is : %d", num, result)
	} else {
		fmt.Printf("%d is Odd, half of it is : %d", num, result)

	}
}
