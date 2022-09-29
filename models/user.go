package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string  `json:"name"`
	Email     string  `json:"email" gorm:"uniqueIndex;type:varchar(255)"`
	Password  []byte  `json:"password"`
	Contact   uint64  `json:"contact"`
	CompanyID uint64  `json:"company_id"`
	Company   Company `gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	// JobsPosted []Job `json:"jobs_posted" gorm:"foreignKey:EmployerID"`
}
