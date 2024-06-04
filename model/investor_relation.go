package model

import (
	"time"

	"gorm.io/gorm"
)

type InvestorRelation struct {
	ID uint `gorm:"primarykey" json:"id"`

	Banner      []*Banner      `json:"banner"`
	Image       string         `gorm:"type:varchar(255)"  json:"image"`
	Title       string         `gorm:"type:varchar(255)"  json:"title"`
	Description string         `gorm:"type:varchar(255)"  json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
