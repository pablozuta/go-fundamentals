package main

import (
	"encoding/json"
	"fmt"
)

type Album struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

func main() {
	albums := []Album{
		{Title: "The Dark Side of the Moon", Artist: "Pink Floyd", Year: 1973},
		{Title: "Thriller", Artist: "Michael Jackson", Year: 1982},
		{Title: "Back in Black", Artist: "AC/DC", Year: 1980},
	}
	b, err := json.Marshal(albums)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(b))
}
