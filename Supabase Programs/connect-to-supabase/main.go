package main

import (
	"encoding/json"
	"fmt"
	"os"

	supa "github.com/nedpals/supabase-go"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	// creamos un cliente supabase
	supabaseUrl := "[URL]"
	supabaseKey := "[KEY]"
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	// usamos el struct Book para guardar los datos obtenidos en un slice llamado 'results'
	var results []Book
	err := supabase.DB.From("books").Select("title, author").Execute(&results)
	if err != nil {
		panic(err)
	}
    
	// mostramos los resultados en formato JSON por consola
	for _, book := range results {
		fmt.Printf("%+v\n", book)
	}
	fmt.Println()

	// escribimos la data a archivo .txt
	fileName := "data.txt"
	fileData, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}
	fmt.Println("Se creo el archivo con datos exitosamente.")
	// el codigo 0644 significa permiso para escribir y leer en un archivo
	err = os.WriteFile(fileName, fileData, 0644)
	if err != nil {
		panic(err)
	}
}
