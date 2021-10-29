package controllers

import (
	"net/http"

	"github.com/Zephiros/amarlinda-api/database"
	"github.com/Zephiros/amarlinda-api/models"
	"github.com/gin-gonic/gin"
)

type CreateAndUpdateProductRequest struct {
	Code            string  `json:"code" binding:"required"`
	Name            string  `json:"name" binding:"required"`
	Description     string  `json:"description"`
	AmountAvailable int64   `json:"amount_available"`
	Price           float32 `json:"price"`
	PricePromotion  float32 `json:"price_promotion"`
	Active          bool    `json:"active"`
	Promotion       bool    `json:"promotion"`
}

// GetProduct ... Get the product by id
// @Summary Get product
// @Description Get product by ID
// @Tags Products
// @Param id path string true "Product ID"
// @Success 200 {object} models.Product
// @Failure 400,401,404 {object} object
// @Router /products/{id} [get]
func GetProduct(c *gin.Context) {
	var product models.Product
	if err := database.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// GetProducts ... Get all products
// @Summary List products
// @Description Get all products
// @Tags Products
// @Success 200 {array} models.Product
// @Failure 401,404 {object} object
// @Router /products [get]
func GetProducts(c *gin.Context) {
	products := []models.Product{}
	query := database.DB.Select("products.*").Group("products.id")
	if err := query.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}

// UpdateProduct ... Update Product
// @Summary Update a product
// @Description Update an existing product based on ID and body parameters
// @Tags Products
// @Accept json
// @Param id path string true "Product ID"
// @Param Product body CreateAndUpdateProductRequest true "Product Data"
// @Success 200 {object} object
// @Failure 400,401,500 {object} object
// @Router /products/{id} [patch]
func UpdateProduct(c *gin.Context) {
	var product models.Product
	if err := database.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := database.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct ... Delete Product
// @Summary Delete a product
// @Description Delete an existing Product by ID
// @Tags Products
// @Accept json
// @Param id path string true "Product ID"
// @Success 200 {object} object
// @Failure 400,401,500 {object} object
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := database.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := database.DB.Where("id = ? ", id).Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
}

// CreateProduct ... Create Product
// @Summary Add a product
// @Description Create new Product based on body parameters
// @Tags Products
// @Accept json
// @Param Product body CreateAndUpdateProductRequest true "Product Data"
// @Success 200 {object} object
// @Failure 400,401,500 {object} object
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	product := models.Product{}
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := database.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}
