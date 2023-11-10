package main

import (
	"social-api/internal/initializers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	initializers.SetupRoutes(app)

	app.Listen(":3000")
}
