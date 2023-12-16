package userUseCases

import (
	userDTO "social-api/internal/models/user/dtos"
	userRepositories "social-api/internal/models/user/repositories"
	"social-api/internal/shared"
)

func FindUserByID(id string) (*userDTO.UserResponse, *shared.AppError) {
	user, err := userRepositories.FindByID(id)
	if err != nil {
		return nil, err
	}

	userResponse := userDTO.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	return &userResponse, nil
}
