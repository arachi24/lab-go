package model

import (
	"time"

	"gorm.io/gorm"
)

type Banner struct {
	ID uint `gorm:"primarykey" json:"id"`

	Image              string         `gorm:"type:varchar(255)"  json:"Image"`
	Link               string         `gorm:"type:varchar(255);not null"  json:"link"`
	Title              string         `gorm:"type:varchar(255)"  json:"title"`
	Description        string         `gorm:"type:varchar(255)"  json:"description"`
	Type               string         `json:"type"`
	HomepageID         *uint          `json:"homepage_id"`
	InvestorRelationID *uint          `json:"investor_id"`
	ProductAreaID      *uint          `json:"product_area_id"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
