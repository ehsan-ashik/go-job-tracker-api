package model

import (
	"gorm.io/gorm"
)

type JobCategory struct {
	gorm.Model
	Name        string  `json:"name" gorm:"type:varchar(255); not null"`
	Description *string `json:"description" gorm:"type:varchar(255); null"`
}
