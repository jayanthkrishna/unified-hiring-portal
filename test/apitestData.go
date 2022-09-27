package test

import (
	"fmt"
	"unified-hiring-portal/database"
	"unified-hiring-portal/models"
)

func TestClientData() {
	clients := []models.Client{
		{
			Secret: "wgsvv34twegv34tgv3q4wte",
			Name:   "naukri",
			Url:    "www.naukri.com",
		},
		{
			Secret: "wgsvv34twegv34tgv3q4wtesrsf",
			Name:   "indeed",
			Url:    "www.indeed.com",
		},
		{
			Secret: "erfbbvsvcegv34tgv3q4wte",
			Name:   "linkedin",
			Url:    "www.linkedin.com",
		},
	}

	database.DB.CreateInBatches(&clients, len(clients))

	res := []models.Client{}

	database.DB.Find(&res)

	for _, i := range res {
		fmt.Printf("Client ID: %v Name: %s Secret: %s\n", i.ID, i.Name, i.Secret)
	}
}
