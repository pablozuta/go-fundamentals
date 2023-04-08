package main

import "fmt"

func main() {
	var size int
	fmt.Print("Ingrese el tamaño de la Pizza (10, 12, 14, 16, 18): ")
	fmt.Scan(&size)

	switch size {
	case 10:
		fmt.Println("el precio de la pizza es $9.990")
	case 12:
		fmt.Println("el precio de la pizza es $11.990")
	case 14:
		fmt.Println("el precio de la pizza es $12.990")
	case 16:
		fmt.Println("el precio de la pizza es $14.990")
	case 18:
		fmt.Println("el precio de la pizza es $18.990")
	default:
		fmt.Println("no ingreso un numero valido")
	}
}