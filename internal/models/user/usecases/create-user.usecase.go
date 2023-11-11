package usecases

import (
	"errors"
	"social-api/config"
	"social-api/internal/entity"
)

func CreateUserUseCase(user *entity.User) error {
	if err := user.ValidateUser(); err != nil {
		return err
	}

	if err := user.EncryptPassword(&user.Password); err != nil {
		return err
	}

	if err := config.DB.Where("email = ?", user.Email).First(&entity.User{}).Error; err == nil {
		return errors.New("email already exists")
	}

	err := config.DB.Create(user).Error
	if err != nil {
		return errors.New("failed to create user")
	}

	return nil
}
