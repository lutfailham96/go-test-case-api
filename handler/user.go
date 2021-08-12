package handler

import (
	"strconv"
	"test-case-api/database"
	"test-case-api/model"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CreateUser new user
func CreateUser(c *fiber.Ctx) error {
	type NewUser struct {
		Email     string `json:"email"`
		Username  string `json:"username"`
		Name      string `json:"name"`
		Address   string `json:"address"`
		Role      string `json:"role"`
		AvatarUrl string `json:"avatar_url"`
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
		Email:     user.Email,
		Username:  user.Username,
		Name:      user.Name,
		Address:   user.Address,
		Role:      user.Role,
		AvatarUrl: user.AvatarUrl,
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
}

func validToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	if uid != n {
		return false
	}

	return true
}

// UpdateUser update user
func UpdateUser(c *fiber.Ctx) error {
	type UpdateUserInput struct {
		Name      string `json:"name"`
		Address   string `json:"address"`
		Role      string `json:"role"`
		AvatarUrl string `json:"avatar_url"`
	}
	var uui UpdateUserInput
	if err := c.BodyParser(&uui); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})
	}

	db := database.DB
	var user model.User

	db.First(&user, id)
	user.Name = uui.Name
	user.Address = uui.Address
	var checkRole = false
	for _, role := range []string{"author", "visitor"} {
		if role == uui.Role {
			checkRole = true
		}
	}
	if !checkRole {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Roles only contains author or visitor", "data": nil})
	}
	user.Role = uui.Role
	user.AvatarUrl = uui.AvatarUrl
	db.Save(&user)

	return c.JSON(fiber.Map{"status": "success", "message": "User successfully updated", "data": user})
}

// GetUser get a user
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var user model.User
	db.Find(&user, id)
	if user.Username == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}

func ChangePassword(c *fiber.Ctx) error {
	db := database.DB

	type changePassword struct {
		Password    string `json:"password"`
		NewPassword string `json:"new_password"`
	}

	var cui changePassword
	if err := c.BodyParser(&cui); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	var user model.User
	userId := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["user_id"]
	db.Find(&user, userId)

	if CheckPasswordHash(cui.Password, user.Password) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Password incorrect", "data": nil})
	}

	hash, err := hashPassword(cui.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
	}

	user.Password = hash
	db.Save(&user)
	return c.JSON(fiber.Map{"status": "success", "message": "Password changed", "data": nil})
}

func GetCurrentUser(c *fiber.Ctx) error {
	userId := uint(c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["user_id"].(float64))
	db := database.DB
	var user model.User
	db.Find(&user, userId)
	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Current user", "current_user": user})
}
