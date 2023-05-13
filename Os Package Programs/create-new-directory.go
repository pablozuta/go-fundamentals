package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Mkdir("NuevaCarpeta", 0755)
	if err != nil {
		fmt.Println("No se pudo crear directorio")
	} else {
		fmt.Println("Directorio creado exitosamente")
	}

}