package models

import (
	"gorm.io/gorm"
)

type Job struct {
	Base
	JobTitle    string      `json:"name" gorm:"not null;type:varchar(255)"`
	Company     string      `json:"company" gorm:"not null;type:varchar(255)"`
	Description string      `json:"description" gorm:"not null;type:varchar(10000)"`
	Position    string      `json:"position"`
	Tags        []Tag       `json:"tags" gorm:"many2many:job_tags;"`
	Applicants  []Applicant `json:"applicants" gorm:"many2many:job_applications;"`
	EmployerID  uint        `json:"employer_id"`
	Employer    User        `gorm:"foreignKey:EmployerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Tag struct {
	gorm.Model
	Name string `json:"tag_name" gorm:"uniqueIndex;type:varchar(255)"`
	Jobs []Job  `json:"jobs" gorm:"many2many:job_tags;"`
}
