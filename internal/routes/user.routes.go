package routes

import (
	userControllers "social-api/internal/models/user/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	userGroup := app.Group("/users")

	userGroup.Get("/", userControllers.GetUser)
	userGroup.Post("/", userControllers.CreateUser)
	userGroup.Put("/:id", userControllers.UpdateUser)
	userGroup.Delete("/:id", userControllers.DeleteUser)
}
