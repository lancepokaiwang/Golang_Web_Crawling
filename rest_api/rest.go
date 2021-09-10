package rest_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// product represents data of merchandise.
type product struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Rating float64 `json:"rating"`
	Url    string  `json:"url"`
}

var products = []product{
	{ID: "1", Name: "A", Price: 1.99, Rating: 4.5, Url: "https://google.com"},
	{ID: "2", Name: "B", Price: 2.99, Rating: 4.7, Url: "https://yahoo.com"},
	{ID: "3", Name: "C", Price: 3.99, Rating: 4.8, Url: "https://apple.com"},
}

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

func NewConnection() error {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.POST("/products", postProducts)

	return router.Run("localhost:8080")
}

// postProducts adds an postProducts from JSON received in the request body.
func postProducts(c *gin.Context) {
	var newProduct product

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	// Add the new album to the slice.
	products = append(products, newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)
}
