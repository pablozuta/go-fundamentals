package main

import "fmt"

func main() {
	// creamos un array con 5 elementos de tipo integer
	arrayUno := [5]int{20, 32, 321, 41, 41}
	// asignamos el valor del primer index del array a una variable
	maximoUno := arrayUno[0]

	// creamos un for loop empezando desde el index 1 que comparara el valor de cada elemento del array con el valor de la variable que empieza en index 0
	// si el valor es mayor se reasignara ese valor a la variable
	for i := 1; i < len(arrayUno); i++ {
		if arrayUno[i] > maximoUno {
			maximoUno = arrayUno[i]
		}
	}
	// mostramos el resultado final por pantalla
	fmt.Printf("El numero mayor en el arrayUno es: %d\n", maximoUno)

	arrayDos := [4]int{4, 63, 32, 111}
	maximoDos := arrayDos[0]

	for i := 1; i < len(arrayDos); i++ {
		if arrayDos[i] > maximoDos {
			maximoDos = arrayDos[i]
		}
	}
	fmt.Printf("El numero mayor en el arrayDos es: %d\n", maximoDos)

}
