package main

import (
	"os"
	"fmt"
)

func main() {
	err := os.Rename("archivo-uno.txt", "archivo-dos.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("archivo renombrado exitosamente.")
	}
}