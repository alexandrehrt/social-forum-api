package main

import (
	"fmt"
	"social-api/config"
	"social-api/internal/initializers"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	config.LoadConfig(".env")
	config.ConnectToDB()
}

func main() {
	app := fiber.New()

	initializers.Routes(app)

	port := viper.Get("WEB_SERVER_PORT")

	app.Listen(fmt.Sprintf(":%s", port))
}
