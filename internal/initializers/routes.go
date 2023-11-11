package initializers

import (
	"social-api/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	routes.UserRoutes(app)
	routes.GroupRoutes(app)
}