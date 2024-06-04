package model

import (
	"time"

	"gorm.io/gorm"
)

type Homepage struct {
	ID uint `gorm:"primarykey" json:"id"`

	// Banner []Banner `gorm:"foreignKey:RefID;references:ID" json:"banner"`
	MainBanner *Banner `json:"main_banner"`
	SubBanner  *Banner `json:"sub_banner"`
	// MainBannerLink string         `gorm:"type:varchar(255);not null" validate:"required"  json:"main_banner_link"`
	// SubBanner *[]Banner `gorm:"foreignKey:RefID;references:ID" json:"sub_banner"`
	// SubBannerLink  string         `gorm:"type:varchar(255)" validate:"required" json:"sub_banner_link"`
	Title       string         `gorm:"type:varchar(255)"  json:"title"`
	SubTitle    string         `gorm:"type:varchar(255)"  json:"sub_title"`
	Description string         `gorm:"type:varchar(255)"  json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
