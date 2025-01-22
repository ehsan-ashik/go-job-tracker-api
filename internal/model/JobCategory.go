package model

type JobCategory struct {
	GormModel
	Name        string  `json:"name" gorm:"type:varchar(255); unique; not null"`
	Description *string `json:"description" gorm:"type:varchar(255); null"`
	Jobs        []Job   `json:"jobs"`
}
