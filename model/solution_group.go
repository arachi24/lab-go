package model

import (
	"time"

	"gorm.io/datatypes"
)

type SolutionGroup struct {
	ID uint `gorm:"primarykey" json:"id"`

	SmartSolutionID uint                           `gorm:"smart_solution_id;index" json:"smart_solution_id"`
	Name            string                         `gorm:"type:varchar(255);not null" validate:"required"  json:"name"`
	Slug            string                         `gorm:"type:varchar(255)"  json:"slug"`
	Solution        datatypes.JSONType[[]Solution] `gorm:"type:text" json:"solutions"`
	CreatedAt       time.Time                      `json:"created_at"`
	UpdatedAt       time.Time                      `json:"updated_at"`
	DeletedAt       *time.Time                     `gorm:"index" json:"deleted_at"`
}

type Solution struct {
	Image          string `json:"image"`
	Link           string `json:"link"`
	IsExternalLink bool   `json:"is_external_link"`
}
