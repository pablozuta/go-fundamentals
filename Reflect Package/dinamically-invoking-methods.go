package main

import (
	"fmt"
	"reflect"
)

type Calculator struct{}

func (c *Calculator) Add(a, b int) int {
	return a + b
}

func main() {
	c := &Calculator{}

	// dinamically invoke the Add method
	method := reflect.ValueOf(c).MethodByName("Add")
	args := []reflect.Value{reflect.ValueOf(5), reflect.ValueOf(3)}
	result := method.Call(args)[0].Int()

	fmt.Println("Resultado de la suma 5 + 8 =",result)
}
