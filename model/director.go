package model

import (
	"time"

	"gorm.io/gorm"
)

type Director struct {
	ID uint `gorm:"primarykey" json:"id"`

	Avatar      string         `gorm:"type:varchar(255);not null" validate:"required"  json:"avatar"`
	Firstname   string         `gorm:"type:varchar(255);not null" validate:"required"  json:"firstname"`
	Lastname    string         `gorm:"type:varchar(255)" validate:"required" json:"lastname"`
	Position    string         `gorm:"type:varchar(255)" validate:"required" json:"position"`
	Description string         `gorm:"type:varchar(255)"  json:"description"`
	IsDirector  bool           `gorm:"column:is_director;default:false" json:"is_director"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
