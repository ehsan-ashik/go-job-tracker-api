package model

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name           string `json:"name" gorm:"type:varchar(255); not null"`
	CareerCiteLink string `json:"career_cite_link" gorm:"type:varchar(255); null"`
	Excitement     *int   `json:"excitement" gorm:"type:integer; null"`
	Jobs           []Job  `json:"jobs"`
}
