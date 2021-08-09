package router

import (
	"test-case-api/handler"
	"test-case-api/middleware"

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
	api.Put("/change-password", middleware.Protected(), handler.ChangePassword)

	// Image
	api.Post("/img", handler.UploadFile)

	// Article
	article := api.Group("/article")
	article.Get("/", middleware.Protected(), handler.GetAllArticles)
	article.Post("/", middleware.Protected(), handler.CreateArticle)
	article.Get("/:id", middleware.Protected(), handler.GetArticle)
	article.Put("/:id", middleware.Protected(), handler.UpdateArticle)

	// Comment
	comment := api.Group("/comment")
	comment.Post("/", middleware.Protected(), handler.CreateComment)
}
