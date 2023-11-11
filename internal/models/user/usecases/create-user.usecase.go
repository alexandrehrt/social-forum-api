package userUseCases

import (
	"errors"
	"social-api/config"
	"social-api/internal/entity"

	"gorm.io/gorm"
)

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	gorm.Model
}

func Create(user *entity.User) (*UserResponse, error) {
	if err := user.ValidateUser(); err != nil {
		return nil, err
	}

	if err := user.EncryptPassword(); err != nil {
		return nil, err
	}

	if userExists := config.DB.Where("email = ?", user.Email).First(&entity.User{}).Error; userExists == nil {
		return nil, errors.New("email already exists")
	}

	if err := config.DB.Create(user).Error; err != nil {
		return nil, errors.New("failed to create user")
	}

	userResponse := UserResponse{
		Username: user.Username,
		Email:    user.Email,
		Model:    user.Model,
	}

	return &userResponse, nil
}
