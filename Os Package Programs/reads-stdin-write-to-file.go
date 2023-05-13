package main

import (
	"fmt"
	"os"
)		

func main() {
	fileName := "output.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}	
	defer file.Close()

	var input string
	fmt.Println("Ingrese una frase (ingrese la palabra 'salir' para abandonar el programa)")
	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println(err)
			break
		}
		_, err = file.WriteString(input + "\n")
		if err != nil {
			fmt.Println(err)
			break
		}
		if input == "salir" {
			break
		}
	}
	fmt.Printf("Texto grabado en el archivo %s\n", fileName)
}