package router

import (
	"github.com/gofiber/fiber/v2"

	"backend-go/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/healthz", controllers.GetHealthz)

	users := app.Group("/users")
	users.Get("/", controllers.GetAllProducts)
	// users.Post("/", controllers.createProduct)
	// users.Put("/:id", controllers.updateProduct)
	// users.Delete("/:id", controllers.deleteProduct)

	orders := app.Group("/orders")
	orders.Get("/", controllers.GetHealthz)
	orders.Post("/", controllers.GetHealthz)
	orders.Put("/:id", controllers.GetHealthz)
	orders.Delete("/:id", controllers.GetHealthz)

	products := app.Group("/products")
	products.Get("/", controllers.GetAllProducts)
	products.Post("/", controllers.CreateProduct)
	products.Put("/:id", controllers.UpdateProduct)
	products.Delete("/:id", controllers.DeleteProduct)

	app.Get("/payments", controllers.GetHealthz)
}
