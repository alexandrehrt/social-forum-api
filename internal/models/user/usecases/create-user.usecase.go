package userUseCases

import (
	"errors"
	"social-api/internal/entity"
	userRepositories "social-api/internal/models/user/repositories"
)

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func Create(user *entity.User) (*UserResponse, error) {
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
		return nil, errors.New("failed to create user")
	}

	userResponse := UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	return &userResponse, nil
}
