package test

import (
	"encoding/json"
	"fmt"
	"log"
	"unified-hiring-portal/database"
	"unified-hiring-portal/models"
)

var users = []models.User{
	{
		Name:     "Steven victor",
		Email:    "steven@gmail.com",
		Password: []byte("password"),
	},
	{
		Name:     "Martin Luther",
		Email:    "luther@gmail.com",
		Password: []byte("password"),
	},
	{
		Name:     "Jayanth Luther",
		Email:    "jayanth@gmail.com",
		Password: []byte("password"),
	},
	{
		Name:     "Krishna Luther",
		Email:    "krishna@gmail.com",
		Password: []byte("password"),
	},
	{
		Name:     "yadav Luther",
		Email:    "yadav@gmail.com",
		Password: []byte("password"),
	},
}

var companies = []models.Company{

	{
		Name: "Apple",
	},
	{
		Name: "Google",
	},
}

func TestDataUser() {

	for i, _ := range companies {
		err := database.DB.Create(&companies[i]).Error

		if err != nil {
			log.Fatal("Cannot seed companies table :", err)
		}

	}

	for i, _ := range users {

		users[i].CompanyID = companies[i%2].ID
		err := database.DB.Create(&users[i]).Error

		if err != nil {
			log.Fatal("Cannot seed users table :", err)
		}

	}

	// var res []models.User

	res_users := []models.User{}

	database.DB.Find(&res_users)

	for i, _ := range res_users {
		database.DB.First(&models.Company{}).Where("id = ?", res_users[i].CompanyID).Take(&res_users[i].Company)
	}
	r, _ := json.Marshal(res_users[3])

	fmt.Println("Result after seeding :", string(r))

}

func TestDataJob() {

	jobs := []models.Job{
		{
			JobTitle:    "FrontEnd Engineer",
			Description: "This is a forntend position",
		},
		{
			JobTitle:    "BackEnd Engineer",
			Description: "This is a Backend position",
		},
		{
			JobTitle:    "Devops Engineer",
			Description: "This is a Devops position",
		},
		{
			JobTitle:    "FrontEnd Engineer",
			Description: "This is a Frontend position",
		},
		{
			JobTitle:    "Full Stack Engineer",
			Description: "This is a Full Stack position",
		},
		{
			JobTitle:    "FrontEnd Engineer",
			Description: "This is a forntend position",
		},
	}

	for i, _ := range jobs {
		jobs[i].EmployerID = users[i%2].ID + 1
		err := database.DB.Create(&jobs[i]).Error

		if err != nil {
			log.Fatal("Cannot seed Jobs table :", err)
		}
	}

	res_users := []models.Job{}

	// res_users := []models.User{}

	// database.DB.Preload("Jobs").Find(&res_users)

	err := database.DB.Preload("Applicants").Find(&res_users).Error

	fmt.Println(err)
	// for i, _ := range res_jobs {
	// 	database.DB.First(&models.User{}).Where("id = ?", res_jobs[i].EmployerID).Take(&res_jobs[i].Employer)
	// }

}
