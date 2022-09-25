package models

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID          uuid.UUID  `json:"client_id" gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at"`
	Name        string     `json:"name" gorm:"uniqueIndex;type:varchar(255)"`
	Url         string     `json:"url" gorm:"uniqueIndex"`
	Description string     `json:"description"`
	// Applicants  []Applicant `json:"applicants"`
	Secret     string `json:"client_secret"`
	ClientUser string `json:"client_user"`
}
