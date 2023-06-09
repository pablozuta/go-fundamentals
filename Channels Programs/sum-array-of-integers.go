// // This program creates two goroutines that each sum half of the array of integers using the sum function. Each goroutine sends its result to the result channel. The main function receives the results from the channel and calculates the total sum.  

package main

import "fmt"

// creamos una funcion que recibira el slice de integers y los sumara
func sum(values[]int, result chan int)  {
	// creamos una variable que guardara el valor de la suma
	sum := 0
	// usamos un for range para hacer la suma
	for _, v := range values {
		sum += v
	}
	// enviamos el resultado de sum por el channel result
	result <- sum
}

func main() {
	// se crea un slice de integers que se usara para la suma
	values := []int {2, 43, 41, 7, 5, 11, 25, 9, 77}
	// se crea una variable de tipo channel que recibira los resultados
	result := make(chan int)
    // mostramos los valores que se enviaran a la funcion sum
	fmt.Println(values[:len(values) / 2])
	fmt.Println(values[len(values) / 2:])

	go sum(values[:len(values) / 2], result) // se envia la mitad inicial del array a la funcion sum
	go sum(values[len(values) / 2:], result) // se envia la mitad final del array a la funcion sum

	// creamos dos variables que recibiran los resultados desde el channel result
	x, y := <-result, <-result
	fmt.Println("Valor de suma 1:",x)
	fmt.Println("Valor de suma 2:",y)
	fmt.Println(x + y)
}
