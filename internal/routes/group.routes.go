package routes

import "github.com/gofiber/fiber/v2"

func GroupRoutes(app *fiber.App) {
	groupGroup := app.Group("/group")

	groupGroup.Get("/", getGroup)
	groupGroup.Post("/", createGroup)
	groupGroup.Put("/:id", updateGroup)
	groupGroup.Delete("/:id", deleteGroup)
}

func getGroup(c *fiber.Ctx) error {
	return c.SendString("Get Group")
}

func createGroup(c *fiber.Ctx) error {
	return c.SendString("Create Group")
}

func updateGroup(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Update Group " + id)
}

func deleteGroup(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Delete Group " + id)
}
