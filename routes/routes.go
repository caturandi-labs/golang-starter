package routes

import (
	"caturandi-labs/golang-starter/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupApiV1(app *fiber.App, handlers *handlers.Handlers) {
	v1 := app.Group("/api/v1")
	SetupUserRoutes(v1, handlers)
}
