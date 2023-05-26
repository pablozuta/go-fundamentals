package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string  `json:"title"`
	Genre  string  `json:"genre"`
	Rating float64 `json:"rating"`
}

func main() {
	movies := []Movie{
		{Title: "The Shawshank Redemption", Genre: "Drama", Rating: 9.3},
		{Title: "The Godfather", Genre: "Crime", Rating: 9.2},
		{Title: "The Dark Knight", Genre: "Action", Rating: 9.0},
	}
	moviesData, err := json.Marshal(movies)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(moviesData))

	// mostrar cada field del slice of structs(no es necesario usar string())
	for _, movie := range movies {
		fmt.Println(movie)
	}
}
