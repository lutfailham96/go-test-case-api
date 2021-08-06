package router

import (
	"api-fiber-gorm/handler"
	"api-fiber-gorm/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware Auth
	api := app.Group("/api", logger.New())

	// Auth login & register
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)
	auth.Post("/register", handler.CreateUser)

	// User
	user := api.Group("/user")
	user.Get("/:id", handler.GetUser)
	user.Put("/:id", middleware.Protected(), handler.UpdateUser)

	// Article
	article := api.Group("/article")
	article.Post("/", middleware.Protected(), handler.CreateArticle)
}
