package main

import "fmt"

func main() {
	var año int
	fmt.Print("Ingrese un año:")
	fmt.Scan(&año)

	switch {
	case año%400 == 0:
		fmt.Println("leap year")
	case año%100 == 0:
		fmt.Println("NOT a leap year")
	case año%4 == 0:
		fmt.Println("leap year")
	default:
		fmt.Println("NOT a leap year")
	}
}