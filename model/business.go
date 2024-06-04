package model

import (
	"time"

	"gorm.io/gorm"
)

type Business struct {
	ID uint `gorm:"primarykey" json:"id"`

	Logo        string         `gorm:"type:varchar(255)"  json:"logo"`
	Image       string         `gorm:"type:varchar(255)"  json:"image"`
	Link        string         `gorm:"type:varchar(255)" validate:"required" json:"link"`
	Title       string         `gorm:"type:varchar(255)"  json:"title"`
	Description string         `gorm:"type:varchar(255)"  json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
