package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola desde golang server")
	})

	fmt.Println("escuchando en puerto :8080")
	http.ListenAndServe(":8080", nil)
}