package test

import (
	"encoding/json"
	"fmt"
	"log"
	"unified-hiring-portal/database"
	"unified-hiring-portal/models"
)

func TestDataUser() {
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

func testDataJob() {

}
