package hrmssrc

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// Connect creates a new client and connects to the mongo database.
// It returns an error if the connection can't be established.
func Connect() error {
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return err
	}

	db := client.Database(nameMongDb)

	mg = MongoDB{
		Client: client,
		Db:     db,
	}
	return nil
}

func GetEmployees(c *gin.Context) []Employee {
	c.IndentedJSON(http.StatusOK, employees)
	return employees
}

func GetEmployee(email string, c *gin.Context) (Employee, error) {
	for _, e := range employees {
		if e.Email == email {
			return e, nil
		}
	}

	return Employee{}, fmt.Errorf("employee not found")
}

func AddEmployee(employee Employee, c *gin.Context) error {
	employees = append(employees, employee)
	return nil
}

func UpdateEmployee(email string, employee Employee, c *gin.Context) error {
	for i, e := range employees {
		if e.Email == email {
			employees[i] = employee
			return nil
		}
	}

	return fmt.Errorf("employee not found")
}

func DeleteEmployee(c *gin.Context, email string) error {
	for i, e := range employees {
		if e.Email == email {
			employees = append(employees[:i], employees[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("employee not found")
}

func Disconnect() error {
	return mg.Client.Disconnect(context.TODO())
}

func hrmssrc() {
	if err := Connect(); err != nil {
		fmt.Println(err)
		return
	}

	router := gin.Default()

	router.GET("/employees", GetEmployees)
	router.GET("/employees/:email", GetEmployee)
	router.POST("/employees", AddEmployee)
	router.PATCH("/employees/:email", UpdateEmployee)
	router.DELETE("/employees/:email", DeleteEmployee)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
	defer Disconnect()

}
