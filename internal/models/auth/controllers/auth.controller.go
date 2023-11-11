package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	body := new(LoginBody)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	// if err := usecases.LoginUseCase(body); err != nil {
	// 	return c.Status(400).JSON(fiber.Map{
	// 		"success": false,
	// 		"message": err.Error(),
	// 	})
	// }

	return c.Status(200).JSON(body)
}
