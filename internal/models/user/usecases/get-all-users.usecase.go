package userUseCases

import (
	"errors"
	"social-api/internal/models/user/dtos"
	userRepositories "social-api/internal/models/user/repositories"
)

func GetAllUsers() ([]*userDTO.UserResponse, error) {
	users, err := userRepositories.FindAll()
	if err != nil {
		return nil, errors.New("failed to get all users")
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
