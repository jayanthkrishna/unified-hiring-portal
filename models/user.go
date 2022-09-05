package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      *string `json:"name"`
	Email     *string `json:"email" gorm:"uniqueIndex;type:varchar(255)"`
	Password  []byte  `json:"password"`
	Contact   uint64  `json:"contact"`
	CompanyID uint64  `gorm:"not null" json:"-"`
	Company   Company `gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Company struct {
	ID   uint64 `json:"company_id" gorm:"primaryKey;autoIncrement:True"`
	Name string `json:"company_name"`
}