package main

import (
	"social-api/config"
	"social-api/internal/entity"
)

func init() {
	config.LoadConfig(".env")
	config.ConnectToDB()
}

func main() {
	config.DB.AutoMigrate(&entity.User{})
	config.DB.AutoMigrate(&entity.Group{})
}
