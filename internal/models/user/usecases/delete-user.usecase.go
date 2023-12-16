package userUseCases

import (
	userRepositories "social-api/internal/models/user/repositories"
	"social-api/internal/shared"
)

func DeleteUser(id string) *shared.AppError {
	user, appError := userRepositories.FindByID(id)
	if appError != nil {
		return appError
	}

	if user == nil {
		return &shared.AppError{
			Message:    "User not found",
			StatusCode: 404,
		}
	}

	appError = userRepositories.Delete(id)
	if appError != nil {
		return appError
	}

	return nil
}
