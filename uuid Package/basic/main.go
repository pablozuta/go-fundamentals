package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Este es un valor unico que nos entrega la libreria uuid:")
	fmt.Println(uuid.New().String())
}
