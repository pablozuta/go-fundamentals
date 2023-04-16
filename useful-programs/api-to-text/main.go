package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main()  {
	// configurar el cliente http
	client := http.Client{}
	url := "https://jsonplaceholder.typicode.com/posts"  

	// request hacia la api
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("error al hacer el request", err)
		os.Exit(1)
	}
	
	defer resp.Body.Close()
	
	// leer el body de la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error al leer el body de la respuesta", err)
		os.Exit(1)
	}
	
	// parsing de los datos JSON
	var data []map [string] interface {}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("error al hacer parsing de JSON data", err)
		os.Exit(1)
	}
	
	// abrir un archivo para escritura
	file, err := os.Create("data-api.txt")
	if err != nil {
		fmt.Println("no se pudo crear archivo", err)
		os.Exit(1)
	}

	defer file.Close()

	// escribir los datos en el archivo
	for _, item := range data {  
		for key, value := range item {  
		fmt.Fprintf(file, "%s: %v\n", key, value)  
		}  
		fmt.Fprintln(file)  
		}  
		  
		fmt.Println("Datos escritos en el archivo exitosamente.")  
}