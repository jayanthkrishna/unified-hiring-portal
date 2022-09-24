package models

import (
	"gorm.io/gorm"
)

type Applicant struct {
	gorm.Model
	Name        string `json:"name"`
	Email       string `json:"email" gorm:"not null;uniqueIndex;type:varchar(255)"`
	JobsApplied []Job  `json:"jobs_applied" gorm:"many2many:job_applications;"`
}
