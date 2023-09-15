package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		num := rand.Intn(100)
		str := strconv.Itoa(num)
		fmt.Fprintln(w, "<h1>Random Number Generator </h1>")
		fmt.Fprintln(w, "<h2>"+str+"</h2>")
	})
	fmt.Println("Server Ready on Port 8080")
	http.ListenAndServe(":8080", nil)
}
