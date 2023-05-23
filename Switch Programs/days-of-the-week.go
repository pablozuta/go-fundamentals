package main

import "fmt"

func main() {
	var diasDeLaSemana int
	fmt.Print("Ingrese un numero entre 1 y 7:")
	fmt.Scan(&diasDeLaSemana)
    
	
	switch diasDeLaSemana {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miercoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sabado")
	case 7:
		fmt.Println("Domingo")
	default:
		fmt.Println("Numero no Valido")
	}
}