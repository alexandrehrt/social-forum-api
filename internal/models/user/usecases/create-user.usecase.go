package userUseCases

import (
	"errors"
	"fmt"
	"social-api/internal/entity"
	"social-api/internal/models/user/dtos"
	userRepositories "social-api/internal/models/user/repositories"
)

func Create(user *entity.User) (*userDTO.UserResponse, error) {
	if err := user.ValidateUser(); err != nil {
		return nil, err
	}

	if err := user.EncryptPassword(); err != nil {
		return nil, err
	}

	err := userRepositories.FindByUsername(user)
	if err == nil {
		return nil, errors.New("username already exists")
	}

	err = userRepositories.FindByEmail(user)
	if err == nil {
		return nil, errors.New("email already exists")
	}

	user, err = userRepositories.Create(user)
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		return nil, errors.New("failed to create user")
	}

	userResponse := userDTO.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	return &userResponse, nil
}
