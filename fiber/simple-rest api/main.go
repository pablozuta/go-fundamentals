package main  
  
import (  
	"fmt"  
	"log"  
	  
	"github.com/gofiber/fiber/v2"  
)  
  
type Book struct {  
	ID int `json:"id"`  
	Title string `json:"title"`  
	Author string `json:"author"`  
}  
  
var books = []Book{  
	{ID: 1, Title: "The Catcher in the Rye", Author: "J.D. Salinger"},  
	{ID: 2, Title: "To Kill a Mockingbird", Author: "Harper Lee"},  
	{ID: 3, Title: "1984", Author: "George Orwell"},  
}  
  
func main() {  
	app := fiber.New()  
	  
	// GET all books  
	app.Get("/books", func(c *fiber.Ctx) error {  
	return c.JSON(books)  
	})  
	  
	// GET a book by ID  
	app.Get("/books/:id", func(c *fiber.Ctx) error {  
		id := c.Params("id")  
		for _, book := range books {  
		if fmt.Sprintf("%d", book.ID) == id {  
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
		books = append(books, book)  
		return c.SendString("Book added")  
	})  
	  
	// DELETE a book by ID  
	app.Delete("/books/:id", func(c *fiber.Ctx) error {  
		id := c.Params("id")  
		for i, book := range books {  
		if fmt.Sprintf("%d", book.ID) == id {  
			books = append(books[:i], books[i+1:]...)  
			return c.SendString("Book deleted")  
		}  
	}  
		return c.Status(404).SendString("Book not found")  
	})  
	  
	// Start server  
	log.Fatal(app.Listen(":3000"))  
} 