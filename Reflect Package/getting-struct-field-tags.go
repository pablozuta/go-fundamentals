package main

import (
	"fmt"
	"reflect"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func main()  {
	libro := Book {"Pedro Paramo", "Juan Rulfo", 1956}
	fmt.Println("Struct Data:", libro)
	t := reflect.TypeOf(libro)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		fmt.Printf("Field: %s, Tag: %s\n", field.Name, tag)
	}
}