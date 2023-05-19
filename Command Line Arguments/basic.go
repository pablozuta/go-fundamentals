package main

import (
	"fmt"
	"os"
)

func main() {
	for _, args := range os.Args {
		fmt.Println(args)
	}
}

// ejecutar
// go run basic.go argumentoUno argumentoDos
// output = argumentoUno argumentoDos