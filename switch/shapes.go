package main

import "fmt"

func main() {
	exit := false

	for exit == false {
		var shape int
	fmt.Print("ingrese numero de lados de la figura (entre 3 y 8) marque 9 si desea salir del programa: ")
	fmt.Scan(&shape)

	switch shape {
	case 3:
		fmt.Println("la figura es un triangulo")
	case 4:
		fmt.Println("la figura es un cuadrado")
	case 5:
		fmt.Println("la figura es un pentagono")
	case 6:
		fmt.Println("la figura es un hexagono")
	case 7:
		fmt.Println("la figura es un heptagono")
	case 8:
		fmt.Println("la figura es un octageno")
	case 9:
		fmt.Println("gracias por probar el programa")
		exit = true
	default:
		fmt.Println("numero ingresado no es valido")

	}
		
	}
	

	
}