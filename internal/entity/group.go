package entity

import (
	"gorm.io/gorm"
	"time"
)

type Group struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name        string         `json:"name" gorm:"unique; not null" validate:"required"`
	Description string         `json:"description" gorm:"not null" validate:"required"`
	Admin       int            `json:"admin"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}
