package controllers

import (
	"social-api/internal/entity"
	"social-api/internal/models/user/usecases"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {

	return c.SendString("Get User")
}

func CreateUser(c *fiber.Ctx) error {
	user := new(entity.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	if err := usecases.CreateUserUseCase(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Update User " + id)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Delete User " + id)
}
