package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobDescription struct {
	gorm.Model
	JobID       uuid.UUID `json:"job_id" gorm:"type:uuid; not null"`
	Description string    `json:"description" gorm:"type:text; not null"`
	Job         *Job      `json:"job"`
}
