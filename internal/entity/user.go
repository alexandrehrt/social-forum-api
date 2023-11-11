package entity

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Username string `json:"username" gorm:"unique; not null" validate:"required"`
	Email    string `json:"email" gorm:"unique; not null" validate:"required,email"`
	Password string `json:"password" gorm:"not null" validate:"required"`
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
	validate := validator.New()

	err := validate.Struct(u)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Field %s failed validation for tag %s with value: %s\n", err.Field(), err.Tag(), err.Param())

		}

		return err
	}

	return nil
}

func (u *User) EncryptPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	return nil
}

func (u *User) ValidatePassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return errors.New("invalid password")
	}

	return nil
}
