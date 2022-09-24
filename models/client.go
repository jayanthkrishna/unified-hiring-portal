package models

type Client struct {
	Base
	Name        string `json:"name" gorm:"uniqueIndex;type:varchar(255)"`
	Url         string `json:"url" gorm:"uniqueIndex"`
	Description string `json:"description"`
	// Applicants  []Applicant `json:"applicants"`
	Secret string `json:"client_secret"`
}
