package model

import "time"

type Product struct {
	ID uint `gorm:"primarykey" json:"id"`

	Image         string         `gorm:"type:varchar(255)"  json:"image"`
	Link          string         `gorm:"type:varchar(255)"  json:"link"`
	Published     bool           `json:"published"`
	IsProduct     bool           `json:"is_product"`
	ProductDetail *ProductDetail `json:"product_detail"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     *time.Time     `gorm:"index" json:"deleted_at"`
}
