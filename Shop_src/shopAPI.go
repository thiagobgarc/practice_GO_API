package shopsrc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

var products = []Product{
	{
		ID:       "1",
		Name:     "Product 1",
		Price:    10.00,
		Quantity: 10,
	},
	{
		ID:       "2",
		Name:     "Product 2",
		Price:    20.00,
		Quantity: 5,
	},
	{
		ID:       "3",
		Name:     "Product 3",
		Price:    30.00,
		Quantity: 3,
	},
}

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

func addProduct(c *gin.Context) {
	var newProduct Product

	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	products = append(products, newProduct)

	c.IndentedJSON(http.StatusCreated, newProduct)
}

func getProductByName(c *gin.Context) {
	name := c.Param("name")

	for _, a := range products {
		if a.Name == name {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range products {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

func shopsrc() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.POST("/products", addProduct)
	router.GET("/products/:name", getProductByName)
	router.GET("/products/:id", getProductByID)
	router.Run("localhost:8080")
}
