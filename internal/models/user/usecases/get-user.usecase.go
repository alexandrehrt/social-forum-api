package userUseCases

import (
	"errors"
	userDTO "social-api/internal/models/user/dtos"
	userRepositories "social-api/internal/models/user/repositories"
)

func GetUser(id string) (*userDTO.UserResponse, error) {
	user, err := userRepositories.FindById(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	userReponse := userDTO.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	return &userReponse, nil
}
