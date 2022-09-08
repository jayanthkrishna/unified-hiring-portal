package models

import (
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	JobTitle    string `json:"name"`
	Company     string `json:"company"`
	Description string `json:"description"`
	Position    string `json:"position"`
	// ApplicantID uint        `gorm:"default:null" json:"-"`
	Applicants []Applicant `json:"applicants" gorm:"many2many:job_applications;"`
	EmployerID uint        `json:"employer_id"`
	Employer   User        `gorm:"foreignKey:EmployerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Applicant struct {
	gorm.Model
	Name        string `json:"name"`
	JobsApplied []Job  `gorm:"many2many:job_applications;"`
}
