package main  
  
import (  
fiber "github.com/gofiber/fiber/v2"  
)  
  
func main() {  
	// crear una instancia de fiber  
	app := fiber.New()  
	  
	// definimos una ruta 
	app.Get("/", func(c *fiber.Ctx) error {  
	return c.SendString("Hola desde FIBER")  
	})  

	// otra ruta
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("About Page")
	})
	  
	// iniciamos el servidor en el puerto 3000 
	err := app.Listen(":3000")  
	if err != nil {  
	panic(err)  
	}  
}