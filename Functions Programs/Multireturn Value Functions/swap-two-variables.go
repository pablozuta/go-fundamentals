package main

import "fmt"

// la funcion devuelve los argumentos de entrada en orden reverso
func swapVar(a, b int) (int, int) {
	return b, a
}

func main() {
	numeroUno := 20
	numeroDos := 30
	fmt.Printf("Antes de Swap numeroUno = %d, numeroDos = %d\n", numeroUno, numeroDos)
	numeroUno, numeroDos = swapVar(numeroUno, numeroDos)
	fmt.Printf("Despues de Swap numeroUno = %d, numeroDos = %d", numeroUno, numeroDos)

}
