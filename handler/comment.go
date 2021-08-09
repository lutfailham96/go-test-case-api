package handler

import (
	"strconv"
	"test-case-api/database"
	"test-case-api/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func CreateComment(c *fiber.Ctx) error {
	type CommentInput struct {
		CommentText string `json:"comment_text"`
	}
	var cii CommentInput
	if err := c.BodyParser(&cii); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create new comment", "data": err})
	}

	db := database.DB
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	userId := uint(c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["user_id"].(float64))
	comment := new(model.Comment)
	comment.CommentText = cii.CommentText
	comment.ArticleID = uint(id)
	comment.UserID = userId
	db.Create(&comment)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Created comment", "data": comment})
}

func DeleteComment(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var comment model.Comment
	db.Find(&comment, id)

	if comment.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No comment found with ID", "data": nil})
	}
	db.Delete(&comment)
	return c.Status(fiber.StatusGone).JSON(fiber.Map{"status": "success", "message": "Comment deleted", "data": nil})
}
