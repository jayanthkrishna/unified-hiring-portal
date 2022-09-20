package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Name string `json:"name" gorm:"uniqueIndex;type:varchar(255)"`
}
