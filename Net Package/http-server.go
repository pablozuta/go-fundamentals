package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error al inicial el servidor HTTP" + err.Error())
	}
	fmt.Println("Servidor iniciado en puerto 8080")
}
