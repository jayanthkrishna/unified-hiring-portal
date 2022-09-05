package database

import (
	"fmt"
	"unified-hiring-portal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

var DB *gorm.DB

func NewConnection(config *Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err

}

func Migrate(db *gorm.DB) error {
	db.Migrator().DropTable(&models.User{})
	err := db.AutoMigrate(&models.User{}, &models.Company{})

	return err
}
