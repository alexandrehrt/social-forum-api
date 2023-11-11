package userRepositories

import (
	"social-api/config"
	"social-api/internal/entity"
)

func Create(user *entity.User) (*entity.User, error) {
	if err := config.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func FindByUsername(user *entity.User) error {
	if err := config.DB.Where("username = ?", user.Username).First(user).Error; err != nil {
		return err
	}

	return nil
}

func FindByEmail(user *entity.User) error {
	if err := config.DB.Where("email = ?", user.Email).First(user).Error; err != nil {
		return err
	}

	return nil
}
