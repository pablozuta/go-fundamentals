package main

import "fmt"

func main() {
	str := "Altered Carbon"
	reversed := ""
	for i := len(str)-1; i >= 0; i-- {
		// el operador += append una letra a la variable string
		reversed += string(str[i])

	}
	fmt.Println("Reversed String = ", reversed)
}
