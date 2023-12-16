package userRepositories

import (
	"net/http"
	"social-api/config"
	"social-api/internal/entity"
	"social-api/internal/shared"
)

func Create(user *entity.User) (*entity.User, error) {
	if err := config.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func FindAll() ([]*entity.User, *shared.AppError) {
	var users []*entity.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, &shared.AppError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return users, nil
}

func FindByID(id string) (*entity.User, *shared.AppError) {
	var user entity.User
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, &shared.AppError{
			Message:    "User not found",
			StatusCode: http.StatusNotFound,
		}
	}

	return &user, nil
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

func Delete(id string) error {
	if err := config.DB.Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
		return err
	}

	return nil
}
