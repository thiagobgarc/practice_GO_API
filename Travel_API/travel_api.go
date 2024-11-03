package travelapi

import (
    "net/http"
    "fmt"

    "github.com/gin-gonic/gin"
)

type travel struct {
	ID      string `json:"id"`
	Country string `json:"country"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}

var api = []Travel{
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
    var newTravel Travel

    if err := c.ShouldBindJSON(&newTravel); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    api = append(api, newTravel)

    c.IndentedJSON(http.StatusCreated, newTravel)
}


