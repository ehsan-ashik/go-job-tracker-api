package model

import (
	"time"

	"github.com/google/uuid"
)

type Job struct {
	GormModel
	ID             uuid.UUID      `json:"id" gorm:"type:uuid"`
	Position       string         `json:"position" gorm:"type:varchar(255); not null"`
	CompanyID      uint           `json:"company_id" gorm:"type:int; not null"`
	Company        Company        `json:"company"`
	Status         string         `json:"status" gorm:"type:varchar(255); default:'Applied'; not null"`
	ApplyDate      time.Time      `json:"apply_date" gorm:"type:date; default:current_timestamp; not null"`
	ResonseDate    *time.Time     `json:"response_date" gorm:"type:date; default:null; null"`
	Remark         *string        `json:"remark" gorm:"type:text; null"`
	Excitement     *int           `json:"excitement" gorm:"type:integer; null"`
	IsReferred     bool           `json:"is_referred" gorm:"type:boolean; not null"`
	ReferredBy     *string        `json:"referred_by" gorm:"type:varchar(255); null"`
	Location       *string        `json:"location" gorm:"type:varchar(255); default:'USA'; null"`
	JobCategoryID  uint           `json:"job_category_id" gorm:"type:int; not null"`
	JobCategory    JobCategory    `json:"job_category"`
	JobDescription JobDescription `json:"job_description"`
	ResumeID       *uint          `json:"resume_id" gorm:"type:int; null"`
	Resume         Resume         `json:"resume"`
}
