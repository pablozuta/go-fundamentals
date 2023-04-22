package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Done bool `json:"done"`
}

var todos = []Todo {
	{ID:1, Title:"Learn Go", Description: "Learn how to use the go programming language", Done: false},
	{ID:2, Title:"Build a REST API", Description: "Build a REST API using Go and Fiber", Done: false},
	{ID:3, Title:"Deploy to production", Description: "Deploy the REST API to a production environment", Done: false},
}

func main()  {
	app := fiber.New()

	// GET all todos
	app.Get("api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	} )

	// GET a todo by ID
	app.Get("api/todos/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString("Bad Request")
		}
		for _, todo := range todos {
			if todo.ID == id {
				return c.JSON(todo)
			}
		}
		return c.Status(404).SendString("Todo NOT Found")
	})

	// start SERVER
	fmt.Println("Server Listening on port 3000")
	log.Fatal(app.Listen(":3000"))
	
}