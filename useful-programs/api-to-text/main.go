// This program makes an HTTP request to the JSONPlaceholder API to get the data, parses the JSON response, and writes the data to a text file named "data.txt". The data is written in key-value pairs, with each item separated by a blank line.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// configurar el cliente HTTP y URL
	client := http.Client{}
	url := "https://jsonplaceholder.typicode.com/posts"

	// request hacia la API usando el cliente HTTP y URL que creamos
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("error al hacer el request", err)
		os.Exit(1)
	}

	// diferimos el cierre del body de la respuesta hasta que termine la ejecucion del programa
	defer resp.Body.Close()

	// leemos el body de la respuesta y lo guardamos en una variable (la data estara en formato JSON)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error al leer el body de la respuesta", err)
		os.Exit(1)
	}

	// parsing de los datos JSON
	// creamos una variable de tipo map y hacemos un unmarshal para "traducir" entre JSON y el mapa de Go
	var data []map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("error al hacer parsing de JSON data", err)
		os.Exit(1)
	}

	// abrimos un archivo para escribir la data
	file, err := os.Create("data-api.txt")
	if err != nil {
		fmt.Println("no se pudo crear archivo", err)
		os.Exit(1)
	}

	defer file.Close()

	// escribimos los datos en el archivo usando un for range y los metodos Fprintf y Fprintln
	for _, item := range data {
		for key, value := range item {
			fmt.Fprintf(file, "%s: %v\n", key, value)
		}
		fmt.Fprintln(file)
	}

	// mostramos un mensaje de operacion exitosa
	fmt.Println("Datos escritos en el archivo exitosamente.")
}
