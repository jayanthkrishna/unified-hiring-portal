package models

import (
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Name        string      `json:"name"`
	Company     string      `json:"company"`
	Description string      `json:"description"`
	Position    string      `json:"position"`
	ApplicantID uint        `gorm:"not null" json:"-"`
	Applicants  []Applicant `json:"applicants" gorm:"foreignKey:ApplicantID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	EmployerID  uint        `gorm:"not null" json:"-"`
	Employer    User        `gorm:"foreignKey:EmployerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Applicant struct {
	gorm.Model
	Name string `json:"name"`
}
