package main

import (
	"fmt"
	"log"
	"os"
	"unified-hiring-portal/database"
	"unified-hiring-portal/routes"
	"unified-hiring-portal/test"

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

	test.TestDataUser()
	test.TestDataJob()

	test.TestDataApplicants()
	test.TestDataApplications()

	// res := []models.User{}

	// database.DB.Preload("JobsPosted.Applicants").Preload("JobsPosted").Find(&res)

	// fmt.Println("User: ", res[1].Name, res[1].Email)

	// for _, i := range res[1].JobsPosted {

	// 	fmt.Println("Job Title : ", i.JobTitle)
	// 	fmt.Println("Applicants")

	// 	for _, j := range i.Applicants {
	// 		fmt.Printf("Applicant ID : %d --Applicant Name : %s. Applicant Email %s\n", j.ID, j.Name, j.Email)
	// 	}

	// }

	// jobs := []models.Job{}

	// database.DB.Find(&jobs)

	// b, err := json.MarshalIndent(jobs[0], "", "  ")

	// fmt.Println("Job with OmitEmpty : ", string(b))

	// test.TestBaseData()

	test.TestClientData()
	go apiServer()

	server()

}

func server() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.GetRoutes(app)

	app.Listen(":8000")

}

func apiServer() {
	api := fiber.New()

	api.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.GetApiRoutes(api)

	api.Listen(":8001")

}
