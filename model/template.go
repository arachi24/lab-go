package model

import (
	"time"

	"gorm.io/datatypes"
)

type Template struct {
	ID uint `gorm:"primarykey" json:"id"`

	Slug      string         `gorm:"type:varchar(255)"  json:"slug"`
	Styles    datatypes.JSON `gorm:"type:text" json:"styles"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt *time.Time     `gorm:"index" json:"deleted_at"`
}
