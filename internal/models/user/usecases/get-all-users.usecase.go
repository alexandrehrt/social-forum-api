package userUseCases

import (
	"social-api/internal/models/user/dtos"
	userRepositories "social-api/internal/models/user/repositories"
	"social-api/internal/shared"
)

func GetAllUsers() ([]*userDTO.UserResponse, *shared.AppError) {
	users, appError := userRepositories.FindAll()
	if appError != nil {
		return nil, appError
	}

	var userResponses []*userDTO.UserResponse
	for _, user := range users {
		userResponse := userDTO.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		}

		userResponses = append(userResponses, &userResponse)
	}

	return userResponses, nil
}
