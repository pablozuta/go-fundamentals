package main

import "fmt"

func isPrime(numero int) bool {
	if numero < 2 {
		return false
	}
	for i := 2; i < numero/2; i++ {
		if numero%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	numero := 6
	resultado := isPrime(numero)
	fmt.Println("Numero es Primo?", resultado)

}
