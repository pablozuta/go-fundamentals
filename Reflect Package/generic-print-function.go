package main

import (
	"fmt"
	"reflect"
)

func Print(value interface{}) {
	// obtain the reflect.Value of the input value
	v := reflect.ValueOf(value)
	// check the kind of the value
	switch v.Kind() {
	case reflect.Struct:
		// iterate over the struct fields
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			// Print the field's type and value
			fmt.Printf("%s %v\n", field.Type(), field.Interface())
		}
	default:
		fmt.Println("Unsopported Type")
	}

}

// crear un struct Videogame con 2 fields
type Videogame struct {
	Titulo string
	Año    int
}

func main() {
	// crear una instancia del Videogame struct
	v := Videogame{"Streets of Rage", 1991}
	// call the Print function con el valor de v
	Print(v)
}
