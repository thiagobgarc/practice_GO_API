package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type rest_api struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var api = []rest_api{
	{
		ID:       "1",
		Name:     "John Doe",
		Email:    "mY6Y5@example.com",
		Password: "password",
	},
}

func getRestAPI(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, api)
}

func addRestAPI(c *gin.Context) {
	var newAPI rest_api

	if err := c.ShouldBindJSON(&newAPI); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	api = append(api, newAPI)

	c.IndentedJSON(http.StatusCreated, newAPI)
}

func getRestAPIs(c *gin.Context) {
	id := c.Param("id")

	api, error := getRestAPIByID(id)

	if error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "api not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, api)
}

func getRestAPIByID(id string) (rest_api, error) {
	for _, a := range api {
		if a.ID == id {
			return a, nil
		}
	}

	return rest_api{}, fmt.Errorf("api not found")
}

func getRestAPIByName(name string) (rest_api, error) {
	for _, a := range api {
		if a.Name == name {
			return a, nil
		}
	}

	return rest_api{}, fmt.Errorf("api not found")
}

func getRestAPIByEmail(email string) (rest_api, error) {
	for _, a := range api {
		if a.Email == email {
			return a, nil
		}
	}

	return rest_api{}, fmt.Errorf("api not found")
}

func toggleRestAPIStatus(c *gin.Context) {
	id := c.Param("id")

	api, error := getRestAPIByID(id)

	if error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "api not found"})
		return
	}

	// api.Completed = !api.Completed

	c.IndentedJSON(http.StatusOK, api)
}

func main() {
	router := gin.Default()
	router.GET("/api", getRestAPI)
	router.POST("/api", addRestAPI)
	router.GET("/api/:id", getRestAPI)
	router.PATCH("/api/:id", toggleRestAPIStatus)
	router.Run("localhost:8080")
}
