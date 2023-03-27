package main

import (
	"fmt"
)

func main() {
	// crear un slice vacio de diferentes tipos
	miSliceCapacidadCinco := make([]int, 5)
	fmt.Println(miSliceCapacidadCinco) // mostrara solo ceros

	miSliceStrings := make ([]string, 4)
	fmt.Println(miSliceStrings) // muestra un slice vacio

	miSliceBooleans := make ([]bool, 3)
	fmt.Println(miSliceBooleans) // muestra valores "false" por cada elemento

	miSliceFloat := make ([]float32, 6)
	fmt.Println(miSliceFloat) // al igual que int mostrara solo ceros

	// crear e inicializar un slice
	miSlice := []int {2, 433, 41, 76, 2666, 43, 655, 109}
	fmt.Println(miSlice)
	fmt.Println(miSlice[2]) // accede a una ubicacion dentro del slice

	// slice de un slice
	miNuevoSlice := miSlice[1:3] // copia la posicion 1 y 2
	fmt.Println(miNuevoSlice)

	// append un elemento nuevo
	miNuevoSlice = append(miNuevoSlice, 333)
	fmt.Println(miNuevoSlice)

	// saber la capacidad de un slice
	fmt.Println(cap(miSlice)) // nos devuelve 8
	
	// la longitud de un slice
	fmt.Println(len(miSlice)) // nos devuelve 8

	// copiar un slice a otro
	sliceUno := [] string {"Aqui todo", "va de mal en peor", "el aguacero", "llego de repente"}
	sliceDos := make([] string, len(sliceUno))
	copy(sliceDos, sliceUno) // copia el contenido de sliceUno a sliceDos
	fmt.Println(sliceUno) 
	fmt.Println(sliceDos) 


}