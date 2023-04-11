package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	supa "github.com/nedpals/supabase-go"
)

type Person struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

func main() {
	// creamos un cliente supabase
	supabaseUrl := "[SUPABASE_URL]"
	supabaseKey := "[SUPABASE_KEY]"
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	// usamos el struct Person para guardar los datos obtenidos en un slice llamado 'results'
	var results []Person
	err := supabase.DB.From("names").Select("user_name, email").Execute(&results)
	if err != nil {
		panic(err)
	}

	fmt.Println(results) // mostramos los datos por consola

	// escribimos la data a archivo .txt
	fileName := "data.txt"
	fileData, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}
	// el codigo 0644 significa permiso para escribir y leer en un archivo
	err = ioutil.WriteFile(fileName, fileData, 0644)
	if err != nil {
		panic(err)
	}
}
