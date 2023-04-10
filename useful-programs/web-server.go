package main

import (
	"fmt"
	"net/http"
	"log"
)

// funcion para manejar un request
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola desde go server")
}

func main() {
	port := "8080"
	http.HandleFunc("/", handler)
	
	log.Printf("server escuchando en puerto %s", port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}