package routes

import "github.com/gofiber/fiber/v2"

func UserRoutes(app *fiber.App) {
	userGroup := app.Group("/user")

	userGroup.Get("/", getUser)
	userGroup.Post("/", createUser)
	userGroup.Put("/:id", updateUser)
	userGroup.Delete("/:id", deleteUser)
}

func getUser(c *fiber.Ctx) error {
	return c.SendString("Get User")
}

func createUser(c *fiber.Ctx) error {
	return c.SendString("Create User")
}

func updateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Update User " + id)
}

func deleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Delete User " + id)
}
