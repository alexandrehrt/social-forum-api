package main

import (
	"social-api/config"
	"social-api/internal/initializers"

	"github.com/gofiber/fiber/v2"
)

func init() {
	config.ConnectToDB()
}

func main() {
	app := fiber.New()

	initializers.SetupRoutes(app)

	app.Listen(":3000")
}
