package main

import (
	"fmt"
)

// la funcion devuelve el quotient y remainder
func divideAndGetQuotientAndRemainder(dividend, divisor int) (int, int) {
	quotient := dividend / divisor  // resultado de la division
	remainder := dividend % divisor // lo que sobra en la division
	return quotient, remainder
}

func main() {
	dividend, divisor := 27, 4
	quotient, remainder := divideAndGetQuotientAndRemainder(dividend, divisor)
	fmt.Printf("The Quotient of the division (%d / %d) is %d and the remainder is %d\n", dividend, divisor, quotient, remainder)

}
