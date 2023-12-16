package userRepositories

import (
	"net/http"
	"social-api/config"
	"social-api/internal/entity"
	"social-api/internal/shared"
)

func Create(user *entity.User) *shared.AppError {
	if err := config.DB.Create(user).Error; err != nil {
		return &shared.AppError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return nil
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

func FindByUsername(username string) (*entity.User, *shared.AppError) {
	var user entity.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, &shared.AppError{
			Message:    "Username not found",
			StatusCode: http.StatusNotFound,
		}
	}

	return &user, nil
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
