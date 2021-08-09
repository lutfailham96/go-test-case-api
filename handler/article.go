package handler

import (
	"test-case-api/database"
	"test-case-api/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func CreateArticle(c *fiber.Ctx) error {
	db := database.DB
	article := new(model.Article)
	role := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["role"]
	userId := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["user_id"]
	if role != "author" {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "No privilege to create article", "data": nil})
	}
	if err := c.BodyParser(article); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create article", "data": err})
	}
	article.UserID = int(userId.(float64))
	db.Create(&article)
	return c.JSON(fiber.Map{"status": "success", "message": "Created product", "data": article})
}
