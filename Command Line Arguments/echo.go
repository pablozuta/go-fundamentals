package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args is a slice of strings that holds all the arguments passed to the program  
	// The first argument is the name of the executable itself  
	// The subsequent arguments are the command-line arguments passed by the user  
	args := os.Args[1:]

	// mostrara cada palabra que escribamos separada por un espacio
	for _, arg := range args {
		fmt.Printf("%s ", arg)
	}
	fmt.Println()
}
// ejecucion go run echo.go aca escribimos palabras separadas
// output = aca escribimos palabras separadas