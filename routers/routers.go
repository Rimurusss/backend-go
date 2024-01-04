package routers

import (
	"hello-go/controllers"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	// Middleware configuration
	// authMiddleWare := middlewares.NewAuthMiddleware()

	// Controller configuration
	healthzController := controllers.NewHealthzController()
	// exampleController := controllers.NewExampleController()
	// leaveController := controllers.NewLeaveController()

	// Router definition
	app.Get("/healthz", healthzController.GetHealthz)
	// app.Get("/example", exampleController.GetExample)
	// app.Get("/example/error", exampleController.ErrorExample)

	// leave := app.Group("/users")
	// leave.Get("/", leaveController.GetLeaves)
	// leave.Post("/", leaveController.GetLeaves)
	// leave.Put("/:id", leaveController.GetLeaves)
	// leave.Delete("/:id", leaveController.GetLeaves)
	
}
