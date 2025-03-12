package controllers

import (
	"net/http"
	"standard_app/models"
	"github.com/gin-gonic/gin"
)

type Product struct {
}

// @Summary Create a new product
// @Description Create a new product with the given information
// @Tags Products
// @Accept  json
// @Produce  json
// @Param product body Product true "Product data"
// @Success 201 {object} models.Product
// @Router /products [post]
func (p *Product) Create(c *gin.Context) {
	var productModel models.Product
	if err := c.ShouldBindJSON(&productModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Save product
	createdProduct, err := productModel.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}

// @Summary Get all products
// @Description Get a list of all products
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Failure 400 {object} map[string]interface{}
// @Router /products [get]
func (p *Product) List(c *gin.Context) {
	var productModel models.Product
	productList, err := productModel.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": productList})
}

// @Summary Get a product by ID
// @Description Fetch a product by its unique ID
// @Tags Products
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} Product
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products/{id} [get]
func (p *Product) FindById(c *gin.Context) {
	var productModel models.Product
	id := c.Param("id")
	product, err := productModel.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}
