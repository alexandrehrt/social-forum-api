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

func GetAllUsers() ([]*UserResponse, error) {
	users, err := userRepositories.FindAll()
	if err != nil {
		return nil, errors.New("failed to get all users")
	}

	var userResponses []*UserResponse
	for _, user := range users {
		userResponse := UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		}

		userResponses = append(userResponses, &userResponse)
	}

	return userResponses, nil
}

func GetUser(id string) (*UserResponse, error) {
	user, err := userRepositories.FindById(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	userReponse := UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	return &userReponse, nil
}
