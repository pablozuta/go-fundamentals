package main

import (
	"fmt"

)
type Integer struct {
	Value int
}

func (i Integer)Factorial() int  {
	if i.Value < 0 {
		return -1
	}
	if i.Value == 0 {
		return 1
	}
	return i.Value * Integer{Value: i.Value - 1}.Factorial()
}

func main() {
	number1 :=Integer{5}
	number2 :=Integer{7}

	fmt.Printf("%v! = %v\n", number1.Value, number1.Factorial())
	fmt.Printf("%v! = %v\n", number2.Value, number2.Factorial())
}