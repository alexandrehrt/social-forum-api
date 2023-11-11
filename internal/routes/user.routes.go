package routes

import (
	"social-api/internal/models/user/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	userGroup := app.Group("/user")

	userGroup.Get("/", controllers.GetUser)
	userGroup.Post("/", controllers.CreateUser)
	userGroup.Put("/:id", controllers.UpdateUser)
	userGroup.Delete("/:id", controllers.DeleteUser)
}
