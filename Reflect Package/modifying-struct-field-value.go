package main

import (
	"fmt"
	"reflect"
)

type Person3 struct {
	Name string
	Age int
}

func main() {
	p := Person3 {"Albert", 43}
	fmt.Println("Struct Original", p)

	v := reflect.ValueOf(&p).Elem() // obtener el reflect value del struct
	ageField := v.FieldByName("Age") // obtener el reflect value del field Age

	if ageField.IsValid() && ageField.CanSet() {
		ageField.SetInt(20)
	}

	fmt.Println(p)

}