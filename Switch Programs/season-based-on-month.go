package main

import "fmt"

func main() {
	var month int
	fmt.Print("Ingrese un mes entre 1 y 12:")
	fmt.Scan(&month)

	switch month {
	case 12, 1, 2:
		fmt.Println("verano")
	case 3, 4, 5:
		fmt.Println("otoño")
	case 6, 7, 8:
		fmt.Println("invierno")
	case 9, 10, 11:
		fmt.Println("primavera")
	default:
		fmt.Println("no ingreso numero entre 1 y 12")
	}
}