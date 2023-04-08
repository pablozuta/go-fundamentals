package main

import (
	"reflect"
	"fmt"
)

func main() {
	var numero int = 23
	fmt.Println(reflect.TypeOf(numero))
	fmt.Println(reflect.ValueOf(numero))
	
	var numeroFloat float32 = 5.655
	fmt.Println(reflect.TypeOf(numeroFloat))
	
	var mentira bool = false
	fmt.Println(reflect.TypeOf(mentira))
	
	var numeroU uint8 = 34
	fmt.Println(reflect.TypeOf(numeroU))


	// arrays 
	var miArray = [] int {23, 56}
	fmt.Println(reflect.TypeOf(miArray))
	
	var miArrayU = [] uint32 {34, 77, 65, 23}
	fmt.Println(reflect.TypeOf(miArrayU))
	fmt.Println(reflect.ValueOf(miArrayU))


}