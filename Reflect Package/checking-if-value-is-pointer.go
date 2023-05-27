package main

import (
	"reflect"
	"fmt"
)

func main()  {
	value := 42
	valuePtr := &value

	fmt.Println(reflect.ValueOf(value).Kind() == reflect.Ptr)
	fmt.Println(reflect.ValueOf(valuePtr).Kind() == reflect.Ptr)
	fmt.Println(reflect.ValueOf(&valuePtr).Kind() == reflect.Ptr)
	fmt.Println(reflect.ValueOf(&valuePtr).Elem().Kind() == reflect.Ptr)
}