package main

import "fmt"

func saludo(n string)string  {
	return "Hola " + n
}

func main() {
	nombre := "Pablo"
	fmt.Println(saludo(nombre))
}