package controllers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Zephiros/amarlinda-api/database"
	"github.com/Zephiros/amarlinda-api/helpers"
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
	pagination := helpers.GeneratePaginationFromRequest(c, models.Product{})
	var product models.Product
	var products []models.Product
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := database.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

	if err := queryBuider.Model(&models.Product{}).Where(product).Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	pagination.Rows = products

	c.JSON(http.StatusOK, pagination)
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

// ImportProduct ... Import Products
// @Summary Import products
// @Description Import products data
// @Tags Products
// @Param file formData file true "Product data"
// @Success 200 {object} object
// @Failure 400,401,404 {object} object
// @Router /products/import [post]
func ImportProduct(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	r := csv.NewReader(file)
	r.Comma = ';'
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for key, record := range records {
		code := record[0]
		name := record[1]
		description := record[2]
		price := record[3]

		if key == 0 {
			continue
		}
		if len(strings.TrimSpace(record[0])) == 0 {
			break
		}
		priceFloat, err := strconv.ParseFloat(strings.Replace(strings.TrimSpace(price), "R$ ", "", -1), 32)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		var product models.Product
		if err := database.DB.Where("code = ?", code).First(&product).Error; err != nil {
			product := models.Product{
				Code:            &code,
				Name:            &name,
				Description:     description,
				AmountAvailable: 0,
				Price:           float32(priceFloat),
				PricePromotion:  0,
				Active:          false,
				Promotion:       false,
			}
			if err := database.DB.Create(&product).Error; err != nil {
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			product.Name = &name
			product.Description = description
			product.Price = float32(priceFloat)
			if err := database.DB.Save(&product).Error; err != nil {
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}
	}

	c.JSON(http.StatusOK, "OK")
}
