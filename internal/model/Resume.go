package model

import (
	"gorm.io/gorm"
)

type Resume struct {
	gorm.Model
	Title  string  `json:"title" gorm:"type:varchar(255); not null"`
	URL    string  `json:"url" gorm:"type:varchar(255); not null"`
	Remark *string `json:"remark" gorm:"type:text; null"`
	Jobs   []Job   `json:"jobs"`
}
