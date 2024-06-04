package model

import (
	"time"

	"gorm.io/gorm"
)

type News struct {
	ID uint `gorm:"primarykey" json:"id"`

	Title       string         `gorm:"type:varchar(255)"  json:"title"`
	SubTitle    string         `gorm:"type:varchar(255)"  json:"sub_title"`
	Description string         `gorm:"type:varchar(255)"  json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
