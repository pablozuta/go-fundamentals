package main

import (
	"fmt"
	"reflect"
)
type Person struct {
	Name string
	Age int
}


func main() {
	p := Person {"Alice", 32}

	// get the reflect.Type and reflect.Value
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	fmt.Println("Type:", t)
	fmt.Println("Value:", v)

	// getting and setting values
	name := v.FieldByName("Name").Interface().(string)
	age := v.FieldByName("Age").Interface().(int)
	fmt.Println(name, age)

	// inspecting struct fields
	numFields := t.NumField()
	for i := 0; i < numFields; i++ {
		field := t.Field(i)
		fmt.Println(field.Name, field.Type)
	}
}