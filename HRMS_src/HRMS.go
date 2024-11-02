package hrmssrc

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Employee struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Salary    float64 `json:"salary"`
	Active    bool    `json:"active"`
}

var employees = []Employee{
	{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "mY6Y5@example.com",
		Salary:    50000.00,
		Active:    true,
	},
	{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "mY6Y4@example.com",
		Salary:    50000.00,
		Active:    true,
	},
	{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "mY6Y3@example.com",
		Salary:    50000.00,
		Active:    true,
	},
	{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "mY6Y2@example.com",
		Salary:    50000.00,
		Active:    false,
	},
	{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "mY6Y1@example.com",
		Salary:    50000.00,
		Active:    true,
	},
}

type MongoDB struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoDB

const nameMongDb = "HRMS"
const mongoURI = "mongodb://localhost:27017" + nameMongDb
