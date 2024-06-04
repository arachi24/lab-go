package model

import (
	"time"

	"gorm.io/gorm"
)

type Location struct {
	ID uint `gorm:"primarykey" json:"id"`

	Name        string         `gorm:"type:varchar(255)"  json:"name"`
	Address     string         `gorm:"type:varchar(255)"  json:"address"`
	Email       string         `gorm:"type:varchar(255)" json:"email"`
	PhoneNumber string         `gorm:"type:varchar(255)" json:"phone_number"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
