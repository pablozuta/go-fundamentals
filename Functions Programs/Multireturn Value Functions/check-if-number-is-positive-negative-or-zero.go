package main

import "fmt"

func checkNumberType(num int) (string, bool) {
	if num > 0 {
		return "positive", true
	} else if num < 0 {

		return "negative", true
	}
	return "zero", false
}

func main() {
	input := 0
	numberType, isNonZero := checkNumberType(input)
	if isNonZero {
		fmt.Printf("%d is a %s number.\n", input, numberType)
	} else {
		fmt.Printf("The number is Zero")

	}
}
