package database

import (
	"database/sql"
	"fmt"
	"unified-hiring-portal/models"

	"github.com/lib/pq"
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

func NewConnection1(url string) (*sql.DB, error) {

	connection, _ := pq.ParseURL(url)
	connection += " sslmode=require"

	return sql.Open("postgres", connection)

}

func Migrate(db *gorm.DB) error {
	// db.Migrator().DropTable(&models.User{}, &models.Company{})
	db.Migrator().DropTable(&models.Client{}, &models.TestBase{}, &models.User{}, &models.Company{}, &models.Tag{}, &models.Job{}, &models.Applicant{}, "job_applications", "job_tags")

	err := db.AutoMigrate(&models.Client{}, &models.TestBase{}, &models.User{}, &models.Company{}, &models.Tag{}, &models.Job{}, &models.Applicant{})
	// err := db.AutoMigrate(&models.User{}, &models.Company{}, &models.Job{}, &models.Applicant{}, &models.Tag{})

	return err
}
