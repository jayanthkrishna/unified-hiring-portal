package test

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
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

func TestDataApplicants() {

	applicants := []models.Applicant{
		{
			Name:  "Jayanth",
			Email: "jayanth@gmail.com",
		},
		{
			Name:  "Krishna",
			Email: "krishna@gmail.com",
		},
		{
			Name:  "Yadav",
			Email: "yadav@gmail.com",
		},
		{
			Name:  "Abdur",
			Email: "abdur@gmail.com",
		},
		{
			Name:  "Rahiman",
			Email: "rahiman@gmail.com",
		},
		{
			Name:  "Maneesh",
			Email: "maneesh@gmail.com",
		},
		{
			Name:  "Thanuja",
			Email: "thanuja@gmail.com",
		},
		{
			Name:  "Hemanth",
			Email: "hemanth@gmail.com",
		},
		{
			Name:  "Ramana",
			Email: "ramana@gmail.com",
		},
		{
			Name:  "Sri hari",
			Email: "srihari@gmail.com",
		},
		{
			Name:  "Sri Krishna",
			Email: "srikrishna@gmail.com",
		},
		{
			Name:  "Karna",
			Email: "karna@gmail.com",
		},
		{
			Name:  "Harsha",
			Email: "harsha@gmail.com",
		},
	}

	database.DB.CreateInBatches(applicants, len(applicants))
}

func TestDataApplications() {
	jobs := []models.Job{}
	applicants := []models.Applicant{}

	database.DB.Select("ID").Find(&jobs)
	database.DB.Select("ID").Find(&applicants)
	type application struct {
		job_id       uint
		applicant_id uint
	}
	for _, i := range jobs {
		for _, j := range applicants {
			if rand.Intn(100) < 60 {
				application_object := application{
					job_id:       i.ID,
					applicant_id: j.ID,
				}
				database.DB.Table("job_applications").Create(&application_object)
			}
		}

	}

	applications := []application{}

	database.DB.Table("job_applications").Take(&applications)

	for _, i := range applications {
		fmt.Printf("JobID : %d ApplicantID: %d\n", i.job_id, i.applicant_id)
	}

}
