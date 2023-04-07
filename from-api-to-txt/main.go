package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// configuracion del HTTP client
	client := http.Client{}
	url := "https://jsonplaceholder.typicode.com/posts"

	// hacer request a API
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error en API request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// leer el body de la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error al leer el body de la respuesta:", err)
		os.Exit(1)
	}

	// parsing de JSON data
	var data []map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error al hacer parsing de JSON data:", err)
		os.Exit(1)
	}

	// abrir un archivo para escribir
	file, err := os.Create("data.txt")
	if err != nil {
		fmt.Println("Error al crear archivo:", err)
		os.Exit(1)
	}
	defer file.Close()

	// escribir data a archivo
	for _, item := range data {
		for key, value := range item {
			fmt.Fprintf(file, "%s: %v\n", key, value)
		}
		fmt.Fprintln(file)
	}

	fmt.Println("Data guardada exitosamente en archivo.")
}
