package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userID := vars["id"]
		fmt.Fprintf(w, "User ID: %s\n", userID)
	})

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}
