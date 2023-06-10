package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// os.Args is a slice of strings that holds all the arguments passed to the program  
	// The first argument is the name of the executable itself  
	// The subsequent arguments are the command-line arguments passed by the user 
	args := os.Args[1:]

	// inicializar una variable para guardar el resultado
	sum := 0
	// se hace un for range loop para sumar los numeros del slice de argumentos args
	for _, arg := range args {
		// se convierte el argumento a integer usando el package strconv
		num, err := strconv.Atoi(arg)
		if err != nil {
			continue
		}
		// se usa la variable sum para juntar los valores
		sum += num
	}
	fmt.Println(sum)

}

// You can run this program from the command-line with any number of integer arguments, like this: $ go run sum.go 1 2 3 4 5  
