package routes

import (
	"caturandi-labs/golang-starter/config"
	"caturandi-labs/golang-starter/handlers"
	"caturandi-labs/golang-starter/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(grp fiber.Router, handlers *handlers.Handlers) {
	conf := config.New()
	useRoute := grp.Group("/user")
	useRoute.Post("/register", handlers.UserRegister)
	useRoute.Post("/login", handlers.UserLogin)

	// Protected Routes
	useRoute.Use(middleware.IsAuthenticated(conf))
	useRoute.Get("/me", handlers.MeQuery)
}