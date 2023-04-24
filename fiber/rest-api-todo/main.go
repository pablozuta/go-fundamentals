package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

var todos = []Todo{
	{ID: 1, Title: "Learn Go", Description: "Learn how to use the go programming language", Done: false},
	{ID: 2, Title: "Build a REST API", Description: "Build a REST API using Go and Fiber", Done: false},
	{ID: 3, Title: "Deploy to production", Description: "Deploy the REST API to a production environment", Done: false},
}

func main() {
	app := fiber.New()

	// GET all todos
	app.Get("api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

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

	// POST a new todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		var todo Todo
		if err := c.BodyParser(&todo); err != nil {
			return c.Status(400).SendString("Bad Request")
		}
		todo.ID = len(todos) + 1
		todos = append(todos, todo)
		return c.SendString("todo added")
	})

	// PUT update by ID
	app.Put("/api/todos/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString("Bad request")
		}
		var updateTodo Todo
		if err := c.BodyParser(&updateTodo); err != nil {
			return c.Status(400).SendString("Bad request")
		}

		for i, todo := range todos {
			if todo.ID == id {
				todo.Title = updateTodo.Title
				todo.Description = updateTodo.Description
				todo.Done = updateTodo.Done
				todos[i] = todo
				return c.SendString("todo updated")
			}
		}
		return c.Status(404).SendString("Todo not found")
	})

	// DELETE todo by ID
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString("Bad request")
		}

		for i, todo := range todos {
			if todo.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.SendString("Todo deleted")
			}
		}
		return c.Status(404).SendString("Todo not found")
	})

	// start SERVER
	fmt.Println("Server Listening on port 3000")
	log.Fatal(app.Listen(":3000"))

}
