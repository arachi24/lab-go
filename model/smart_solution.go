package model

import (
	"time"
)

type SmartSolution struct {
	ID uint `gorm:"primarykey" json:"id"`

	Banner         string          `gorm:"type:varchar(255);not null" validate:"required"  json:"banner"`
	SolutionGroups []SolutionGroup `json:"solution_groups"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	DeletedAt      *time.Time      `gorm:"index" json:"deleted_at"`
}
