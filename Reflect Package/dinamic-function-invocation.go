// In this example, we have a function Add that takes two integer arguments and returns their sum. Using reflection, we obtain the reflect.Value of the function. We then create a slice of reflect.Value containing the arguments for the function call. Finally, we invoke the function dynamically using Call and retrieve the result.  

package main

import (
	"fmt"
	"reflect"
)

// una funcion que suma dos integers
func Add(a, b int)int  {
	return a + b
}

func main() {
	addFunc := reflect.ValueOf(Add)

	args := []reflect.Value {
		reflect.ValueOf(5),
		reflect.ValueOf(3),
	}

	result := addFunc.Call(args)[0].Int()
	fmt.Println(result)
}