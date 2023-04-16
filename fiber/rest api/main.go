package main

import (
	"fmt"
	
	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
}

var books = []Book {
	{ID: 1, Title: "The Catcher in the Rye", Author: "J.D. Salinger"},
	{ID: 2, Title: "To Kill a Mockingbird", Author: "Harper Lee"},
	{ID: 3, Title: "1984", Author: "George Orwell"},

}

func main()  {
	app := fiber.New()

	// get all books
	app.Get("/books", func(c *fiber.Ctx) error {
		return c.JSON(books)
	})

	app.Listen(":3000")
	fmt.Println("escuchando en puerto 3000")
}