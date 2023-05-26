package main

import (
	"fmt"
	"reflect"
)

type Persona struct {
	Nombre string
	Edad   int
}

func main() {
	t := reflect.TypeOf(Persona{})
	v := reflect.New(t).Elem()

	v.FieldByName("Nombre").SetString("Edgar")
	v.FieldByName("Edad").SetInt(45)

	p := v.Interface().(Persona)
	// mostrar los valores
	fmt.Println(p.Nombre)
	fmt.Println(p.Edad)
}
