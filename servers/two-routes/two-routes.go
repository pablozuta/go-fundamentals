package main

import (
	"fmt"
	"net/http"
)

func main() {
	// se definen rutas
	http.HandleFunc("/", homeHandler) // ruta y nombre de funcion
	http.HandleFunc("/about", aboutHandler)

	// serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	// iniciar servidor
	fmt.Println("se inicia servidor en puerto: 8080")
	http.ListenAndServe(":8080", nil)
}
func homeHandler(w http.ResponseWriter, r *http.Request) {  
	fmt.Fprintf(w, "Welcome to the home page!")  
	}  
	  
	func aboutHandler(w http.ResponseWriter, r *http.Request) {  
	fmt.Fprintf(w, "This is the about page!")  
} 