package userUseCases

import (
	"net/http"
	"social-api/internal/entity"
	userRepositories "social-api/internal/models/user/repositories"
	"social-api/internal/shared"
)

func Create(user *entity.User) *shared.AppError {
	if err := user.ValidateUser(); err != nil {
		return &shared.AppError{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		}
	}

	if err := user.EncryptPassword(); err != nil {
		return &shared.AppError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	userResponse, appError := userRepositories.FindByUsername(user.Username)
	if appError != nil && appError.StatusCode != http.StatusNotFound {
		return appError
	}

	if userResponse != nil && userResponse.Username != "" {
		return &shared.AppError{
			Message:    "username already exists",
			StatusCode: http.StatusBadRequest,
		}
	}

	userResponse, appError = userRepositories.FindByEmail(user.Email)
	if appError != nil && appError.StatusCode != http.StatusNotFound {
		return appError
	}

	if userResponse != nil && userResponse.Email != "" {
		return &shared.AppError{
			Message:    "e-mail already exists",
			StatusCode: http.StatusBadRequest,
		}
	}

	appError = userRepositories.Create(user)
	if appError != nil {
		return &shared.AppError{
			Message:    appError.Message,
			StatusCode: appError.StatusCode,
		}
	}

	return nil
}
