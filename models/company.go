package models

type Company struct {
	ID   uint64 `json:"company_id" gorm:"primaryKey;autoIncrement:True"`
	Name string `json:"company_name"`
	// Users []User `json:"company_users"`
}
