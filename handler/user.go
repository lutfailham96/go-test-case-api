package handler

import (
	"api-fiber-gorm/database"
	"api-fiber-gorm/model"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CreateUser new user
func CreateUser(c *fiber.Ctx) error {
	type NewUser struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Name     string `json:"name"`
		Address  string `json:"address"`
		Role     string `json:"role"`
	}

	db := database.DB
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})

	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
	}

	var checkRole = false
	for _, role := range []string{"author", "visitor"} {
		if role == user.Role {
			checkRole = true
		}
	}
	if !checkRole {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Roles only contains author or visitor", "data": err})
	}

	user.Password = hash
	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	newUser := NewUser{
		Email:    user.Email,
		Username: user.Username,
		Name:     user.Name,
		Address:  user.Address,
		Role:     user.Role,
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
}