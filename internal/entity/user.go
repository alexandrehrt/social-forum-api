package entity

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
	gorm.Model
}

func NewUser(username, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Username: username,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u *User) ValidateUser() error {
	if u.Username == "" {
		return errors.New("username is empty")
	}

	if u.Email == "" {
		return errors.New("email is empty")
	}

	return nil
}

func (u *User) EncryptPassword(password *string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	return nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
