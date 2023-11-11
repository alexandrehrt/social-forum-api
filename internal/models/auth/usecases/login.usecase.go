package usecases

import (
	"errors"
	"fmt"
	"social-api/config"
	"social-api/internal/entity"
)

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginUseCase(body LoginBody) error {
	user := config.DB.Where("username = ?", body.Username).First(&entity.User{})
	if user.Error != nil {
		return errors.New("user not found")
	}
	fmt.Println(user)
	return nil
}
