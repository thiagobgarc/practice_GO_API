package travelapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type travel struct {
	ID      string `json:"id"`
	Country string `json:"country"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}

var api = []travel{
	{
		ID:      "1",
		Country: "United States",
		Name:    "John Doe",
		Email:   "mY6Y5@example.com",
	},
	{
		ID:      "2",
		Country: "United States",
		Name:    "Jane Doe",
		Email:   "mY6Y4@example.com",
	},
	{
		ID:      "3",
		Country: "United States",
		Name:    "John Doe",
		Email:   "mY6Y3@example.com",
	},
	{
		ID:      "4",
		Country: "United States",
		Name:    "Jane Doe",
		Email:   "mY6Y2@example.com",
	},
	{
		ID:      "5",
		Country: "United States",
		Name:    "John Doe",
		Email:   "mY6Y1@example.com",
	},
	{
		ID:      "6",
		Country: "United States",
		Name:    "Jane Doe",
		Email:   "mY6Y0@example.com",
	},
	{
		ID:      "7",
		Country: "United States",
		Name:    "John Doe",
		Email:   "mY6Y9@example.com",
	},
	{
		ID:      "8",
		Country: "United States",
		Name:    "Jane Doe",
		Email:   "mY6Y8@example.com",
	},
	{
		ID:      "9",
		Country: "United States",
		Name:    "John Doe",
		Email:   "mY6Y7@example.com",
	},
}

func getTravel(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, api)
}

func getTravelEmail(c *gin.Context) {
	email := c.Param("email")

	for _, a := range api {
		if a.Email == email {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Email not found"})
}

func getTravelID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range api {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.h{"message": "ID not found"})
}

func addTravel(c *gin.Context) {
	var newTravel travel

	if err := c.ShouldBindJSON(&newTravel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	api = append(api, newTravel)

	c.IndentedJSON(http.StatusCreated, newTravel)
}

func addCountry(c *gin.Context) {
	var newCountry travel

	if err := c.ShouldBindJSON(&newCountry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	api = append(api, newCountry)

	c.IndentedJSON(http.StatusCreated, newCountry)
}

func addName(c *gin.Context) {
	var newName travel

	if err := c.ShouldBindJSON(&newName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	api = append(api, newName)

	c.IndentedJSON(http.StatusCreated, newName)
}

func travelapi() {
	router := gin.Default()
	router.GET("/travel", getTravel)
	router.GET("/travel/:email", getTravelEmail)
	router.GET("/travel/:id", getTravelID)
	router.POST("/travel", addTravel)
	router.POST("/country", addCountry)
	router.POST("/name", addName)
	router.Run("localhost:8080")
}
