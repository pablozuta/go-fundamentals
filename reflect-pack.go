package main

import (
	"reflect"
	"fmt"
)

func main() {
	var numero int = 23
	fmt.Println(reflect.TypeOf(numero))
	
	var numeroFloat float32 = 5.655
	fmt.Println(reflect.TypeOf(numeroFloat))
	
	var mentira bool = false
	fmt.Println(reflect.TypeOf(mentira))

}