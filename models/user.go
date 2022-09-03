package models

type User struct {
	ID       uint64  `json:"id" gorm:"primaryKey;autoIncrement:True"`
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password []byte  `json:"password"`
	Contact  uint64  `json:"contact"`
	Company  Company `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Company struct {
	ID   uint    `json:"company_id" gorm:"primaryKey;autoIncrement:True"`
	Name *string `json:"company_name"`
}
