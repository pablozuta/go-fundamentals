package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Home Page </h1>")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>This is the about page</h1>")
	})
	
	fmt.Println("escuchando en puerto :8080")
	http.ListenAndServe(":8080", nil)

}