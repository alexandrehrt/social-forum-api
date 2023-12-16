package userRepositories

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"social-api/config"
	"social-api/internal/entity"
	"social-api/internal/shared"
	"strings"
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &shared.AppError{
				Message:    "Username not found",
				StatusCode: http.StatusNotFound,
			}
		}

		return nil, &shared.AppError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &user, nil
}

func FindByEmail(email string) (*entity.User, *shared.AppError) {
	var user entity.User
	if err := config.DB.Raw("SELECT * FROM users WHERE email = ?", email).Scan(&user).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, &shared.AppError{
				Message:    "E-mail not found",
				StatusCode: http.StatusNotFound,
			}
		}

		return nil, &shared.AppError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &user, nil
}

func Delete(id string) *shared.AppError {
	result := config.DB.Where("id = ?", id).Delete(&entity.User{})
	if result.Error != nil {
		return &shared.AppError{
			Message:    result.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return nil
}
