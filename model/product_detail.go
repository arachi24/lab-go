package model

import (
	"time"

	"gorm.io/datatypes"
)

type ProductDetail struct {
	ID uint `gorm:"primarykey" json:"id"`

	ProductID uint `gorm:"product_id;index" json:"product_id"`

	Banner         string                                   `gorm:"type:varchar(255)"  json:"banner"`
	Title          string                                   `gorm:"type:varchar(255)"  json:"title"`
	Description    string                                   `gorm:"type:varchar(255)" json:"description"`
	FactSheetUrl   string                                   `gorm:"type:varchar(255)" json:"factsheet_url"`
	HighlightImage string                                   `gorm:"type:varchar(255)" json:"highlight_image"`
	KeyImage       string                                   `gorm:"type:varchar(255)" json:"key_image"`
	FeatureImage   datatypes.JSONType[[]string]             `gorm:"type:text" json:"feature_images"`
	Specifications datatypes.JSONType[[]SpecificationsType] `gorm:"type:text" json:"specifications"`
	Operation      datatypes.JSONType[Operation]            `gorm:"type:text" json:"operation"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}

type SpecificationsType struct {
	Group string               `json:"group"`
	List  []SpecificationsList `json:"list"`
}

type SpecificationsList struct {
	ContentType string `json:"content_type"`
	Label       string `json:"label"`
	Text        string `json:"text"`
}

type Operation struct {
	Published bool                         `json:"published"`
	VideoUrl  string                       `json:"video_url"`
	Image     datatypes.JSONType[[]string] `gorm:"type:text" json:"images"`
}
