package main

import (
	"fmt"
	"log"
	"os"
	"unified-hiring-portal/database"
	"unified-hiring-portal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
	database.DB, err = database.NewConnection(config)
	fmt.Println(config)
	if err != nil {
		log.Fatal("Could not load the database")
	}

	err = database.Migrate(database.DB)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	app := fiber.New()
	api := fiber.New()
	// database.Connect()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	api.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.GetApiRoutes(api)

	routes.GetRoutes(app)

	app.Listen(":8000")
	api.Listen(":8001")
}
