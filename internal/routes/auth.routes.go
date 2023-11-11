package routes

import (
	"social-api/internal/models/auth/controllers"

	"github.com/gofiber/fiber/v2"
)

func Login(app *fiber.App) {
	app.Post("/login", controllers.Login)
}
