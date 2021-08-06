package main

import (
	"api-fiber-gorm/database"
	"api-fiber-gorm/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	database.ConnectDB()
	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
	//defer database.DB.Close()
}
