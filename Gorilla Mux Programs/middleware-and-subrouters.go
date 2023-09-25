package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the main page")
	})

	// subrouter for the /admin	path
	adminRouter := r.PathPrefix("/admin").Subrouter()
	adminRouter.Use(adminMiddleware) // Applying middleware to the subrouter

	adminRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the ADMIN DASHBOARD")
	})

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}

func adminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// perform admin specific checks or actions here
		fmt.Println("Admin Middleware Executed")
		next.ServeHTTP(w, r)
	})
}
