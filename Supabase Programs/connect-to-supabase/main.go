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
	supabaseUrl := "https://nzqszurmnlwplwetnhuz.supabase.co"
	supabaseKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Im56cXN6dXJtbmx3cGx3ZXRuaHV6Iiwicm9sZSI6ImFub24iLCJpYXQiOjE2NzYwODk5NjEsImV4cCI6MTk5MTY2NTk2MX0.CGebwlVumpYn8chKzJi3u64B__SISfW1948cvfJltW0"
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
