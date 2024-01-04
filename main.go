package main

import (
	"github.com/gofiber/fiber/v2"

	"backend-go/database"
	"backend-go/router"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Connect to the Database
	database.ConnectDB()

	// Send a string back for GET calls to the endpoint "/"
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API is UP!")
		return err
	})
	router.SetupRoutes(app)
	// Listen on PORT 3000
	app.Listen(":3000")
}
