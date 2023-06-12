package main

import "fmt"

func main() {
	// tenemos un slice de numeros algunos que se repiten
	numeros := []int {11, 22, 33, 44, 22, 36, 77, 44}
	// creamos un mapa de key/value tipo integer (los valores se inician a su zero value)
	mapaNumeros := make(map[int]int)
	// usamos los elementos del slice como key en el mapa(los mapas no permiten valores repetidos)
	for _, elemento := range numeros {
		mapaNumeros[elemento]++ // incrementamos los valores de key por cada elemento del slice
	}

	// mostramos los keys y valores del mapa (cada vez que los mostremos el orden sera random ya que los elementos en los mapas no estan ordenados de una manera secuencial)
    for key, value := range mapaNumeros {
		fmt.Println(key, value)
	}

	// inicializamos una variable para usarla como contador de los valores del mapa que sean mayor a 1 
	numerosDuplicados := 0
	// usamos un for range loop y dentro un if statement que sumara un valor a la variable por cada valor del mapa que sea
	for _, count := range mapaNumeros {
		if count > 1 {
			numerosDuplicados ++
		}
	}
	// mostramos los resultados
	fmt.Printf("Hay %d numeros duplicados en el mapa", numerosDuplicados)

}