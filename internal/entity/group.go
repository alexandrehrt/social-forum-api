package entity

import (
	"gorm.io/gorm"
)

type Group struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Admin       string `json:"admin"`
	gorm.Model
}

func NewGroup(name, description, admin string) (*Group, error) {
	return &Group{
		Name:        name,
		Description: description,
		Admin:       admin,
	}, nil
}
