package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID     *int   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	// Create books with pointer to ID
	book1 := Book{ID: new(int), Title: "The Catcher in the Rye", Author: "J.D. Salinger"}
	book2 := Book{ID: new(int), Title: "To Kill a Mockingbird", Author: "Harper Lee"}
	book3 := Book{ID: new(int), Title: "1984", Author: "George Orwell"}

	// Assign integer values to book IDs
	*book1.ID = 1
	*book2.ID = 2
	*book3.ID = 3

	// Add books to the list
	books := []Book{book1, book2, book3}

	app := fiber.New()

	// get all books
	app.Get("/books", func(c *fiber.Ctx) error {
		return c.JSON(books)
	})

	// GET a book by ID
	app.Get("/books/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		fmt.Println("Requested ID:", id)
		for _, book := range books {
			fmt.Println("Book ID:", *book.ID)
			if fmt.Sprintf("%d", *book.ID) == id {
				return c.JSON(book)
			}
		}
		return c.Status(404).SendString("Book not found")
	})

	// POST a new book
	app.Post("/books", func(c *fiber.Ctx) error {
		var book Book
		if err := c.BodyParser(&book); err != nil {
			return c.Status(400).SendString("Bad request")
		}
		// Generate a new unique ID for the book
		maxID := 0
		for _, b := range books {
			if *b.ID > maxID {
				maxID = *b.ID
			}
		}
		newID := maxID + 1
		book.ID = &newID
		books = append(books, book)
		return c.SendString("Book added")
	})

	// DELETE a book by ID
	app.Delete("/books/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, book := range books {
			if fmt.Sprintf("%d", book.ID) == id {
				books = append(books[:i], books[i+1:]...)
				return c.SendString("Book deleted.")
			}
		}
		return c.Status(404).SendString("Book Not Found.")
	})

	// start SERVER
	log.Fatal(app.Listen(":3000"))
	fmt.Println("escuchando en puerto 3000")
}

/*
{
    "title": "Pride and Prejudice",
    "author": "Jane Austen"
}

{
    "title": "The Hobbit",
    "author": "J.R.R. Tolkien"
}

{
    "title": "The Hitchhiker's Guide to the Galaxy",
    "author": "Douglas Adams"
}

{
    "title": "Brave New World",
    "author": "Aldous Huxley"
}

{
    "title": "The Lord of the Rings",
    "author": "J.R.R. Tolkien"
}

{
    "title": "The Diary of a Young Girl",
    "author": "Anne Frank"
}

{
    "title": "The Great Expectations",
    "author": "Charles Dickens"
}

{
    "title": "The Little Prince",
    "author": "Antoine de Saint-Exupéry"
}

{
    "title": "The Adventures of Sherlock Holmes",
    "author": "Arthur Conan Doyle"
}

{
    "title": "Frankenstein",
    "author": "Mary Shelley"
}

*/
