package main

import (
	"fmt"
	"reflect"
)

type Persona1 struct {
	Name string
	Age  int
}

func main() {
	var x interface{}
	x = 42

	// check if x is of type Int
	if reflect.TypeOf(x).Kind() == reflect.Int {
		fmt.Println("Numero es Int")
	} else {
		fmt.Println("Numero NO es Int")
	}

	// check if x is of type String
	if reflect.TypeOf(x).Kind() == reflect.String {
		fmt.Println("Valor es String")
	} else {
		fmt.Println("Valor NO es String")
	}

	// check if x is of type Persona1
	if reflect.TypeOf(x).AssignableTo(reflect.TypeOf(Persona1{})) {
		fmt.Println("X can be assigned to a Persona1 struct")
	} else {
		fmt.Println("X can NOT be assigned to a Persona1 struct")
	}

}
