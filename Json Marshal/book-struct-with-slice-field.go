// In this example, we define a Book struct with four fields: Title, Author, Price, and Chapters. The Chapters field is a slice of strings representing the chapters of the book. We create a new Book instance with a title, author, price, and chapters, and then use the json.Marshal function to serialize this object into a JSON-encoded byte slice. Finally, we print the JSON-encoded string to the console.
package main

import (
	"encoding/json"
	"fmt"
)

type book struct {
	title    string   `json:"title"`
	author   string   `json:"author"`
	price    float64  `json:"price"`
	chapters []string `json:"chapters"`
}

func main() {
	bookLord := book{
		title:  "The Lord of the Rings",
		author: "J.R.R. Tolkien",
		price:  19.99,
		chapters: []string{
			"Prologue",
			"Concerning Hobbits",
			"The Shadow of the Past",
			"A Short Cut to Mushrooms",
		},
	}
	jsonData, err := json.Marshal(bookLord)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(string(jsonData))

	fmt.Println(bookLord.title)
	fmt.Println(bookLord.author)
	fmt.Println(bookLord.price)
	fmt.Println("Chapters:")

	for _, chapter := range bookLord.chapters {
		fmt.Println("-" + chapter)
	}

}
