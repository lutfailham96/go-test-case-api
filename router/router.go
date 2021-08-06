package router

import (
	"api-fiber-gorm/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware Auth
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// Auth login & register
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)
	auth.Post("/register", handler.CreateUser)
}
