package handler

import (
	"test-case-api/database"
	"test-case-api/model"

	"github.com/gofiber/fiber/v2"
)

func CreateComment(c *fiber.Ctx) error {
	db := database.DB
	comment := new(model.Comment)
	if err := c.BodyParser(comment); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create new comment", "data": err})
	}
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