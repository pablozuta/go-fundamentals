package main

import "fmt"

// funcion suma que toma dos argumentos de tipo integer y devuelve un resultado tambien de tipo integer
func suma(x, y int)int  {
	return x + y
}

func main() {
	// declaramos dos variables con integers que pasaremos como parametros a la funcion suma
	x := 10
	y := 2
	// llamamos la funcion guardando el resultado en la variable result
	result := suma(x, y)
	// mostramos el resultado
	fmt.Println("Resultado suma es:", result)
}