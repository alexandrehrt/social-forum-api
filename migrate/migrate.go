package main

import (
	"social-api/config"
	"social-api/internal/entity"
)

func init() {
	config.ConnectToDB()
}

func main() {
	config.DB.AutoMigrate(&entity.User{})
}
