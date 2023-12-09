package userUseCases

import (
	"errors"
	userRepositories "social-api/internal/models/user/repositories"
)

func DeleteUser(id string) error {
	err := userRepositories.Delete(id)
	if err != nil {
		return errors.New("failed to delete user")
	}

	return nil
}
