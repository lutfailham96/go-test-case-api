package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"test-case-api/database"
	"test-case-api/model"
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create article", "data": err})
	}
	article.UserID = uint(userId.(float64))
	db.Create(&article)
	return c.JSON(fiber.Map{"status": "success", "message": "Created article", "data": article})
}

func GetAllArticles(c *fiber.Ctx) error {
	db := database.DB
	var articles []model.Article
	db.Preload("Comments").Find(&articles)
	return c.JSON(fiber.Map{"status": "success", "message": "All articles", "data": articles})
}

func GetArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var article model.Article
	db.Preload("Comments").Find(&article, id)
	if article.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No article found with ID", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Article found", "data": article})
}

//func UpdateArticle(c *fiber.Ctx) error {
//	type UpdateArticleInput struct {
//		Title            string `json:"title"`
//		Content          string `json:"content"`
//		FeaturedImageUrl string `json:"featured_image_url"`
//	}
//	var uai UpdateArticleInput
//	if err := c.BodyParser(&uai); err != nil {
//		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
//	}
//
//	db := database.DB
//	var article model.Article
//
//	id := c.Params("id")
//	db.First(&article, id)
//
//	article.Title = uai.Title
//	article.Content = uai.Content
//	article.FeaturedImageUrl = uai.FeaturedImageUrl
//
//	db.Save(&article)
//	return c.JSON(fiber.Map{"status": "success", "message": "Article successfully updated", "data": article})
//}
