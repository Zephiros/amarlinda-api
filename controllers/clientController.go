package controllers

import (
	"net/http"

	"github.com/Zephiros/amarlinda-api/database"
	"github.com/Zephiros/amarlinda-api/helpers"
	"github.com/Zephiros/amarlinda-api/models"
	"github.com/gin-gonic/gin"
)

type CreateAndUpdateClientRequest struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone"`
}

// GetClient ... Get the client by id
// @Summary Get client
// @Description Get client by ID
// @Tags Clients
// @Param id path string true "Client ID"
// @Success 200 {object} models.Client
// @Failure 400,401,404 {object} object
// @Router /clients/{id} [get]
func GetClient(c *gin.Context) {
	var client models.Client
	if err := database.DB.Where("id = ?", c.Param("id")).First(&client).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, client)
}

// GetClients ... Get all clients
// @Summary List clients
// @Description Get all clients
// @Tags Clients
// @Success 200 {array} models.Client
// @Failure 401,404 {object} object
// @Router /clients [get]
func GetClients(c *gin.Context) {
	pagination := helpers.GeneratePaginationFromRequest(c)
	var client models.Client
	var clients []models.Client
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := database.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

	if err := queryBuider.Model(&models.Client{}).Where(client).Find(&clients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	pagination.Rows = clients

	c.JSON(http.StatusOK, pagination)
}

// UpdateClient ... Update Client
// @Summary Update a client
// @Description Update an existing client based on ID and body parameters
// @Tags Clients
// @Accept json
// @Param id path string true "Client ID"
// @Param Client body CreateAndUpdateClientRequest true "Client Data"
// @Success 200 {object} object
// @Failure 400,401,500 {object} object
// @Router /clients/{id} [patch]
func UpdateClient(c *gin.Context) {
	var client models.Client
	if err := database.DB.Where("id = ?", c.Param("id")).First(&client).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.BindJSON(&client); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := database.DB.Save(&client).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, client)
}

// DeleteClient ... Delete Client
// @Summary Delete a client
// @Description Delete an existing client by ID
// @Tags Clients
// @Accept json
// @Param id path string true "Client ID"
// @Success 200 {object} object
// @Failure 400,401,500 {object} object
// @Router /clients/{id} [delete]
func DeleteClient(c *gin.Context) {
	id := c.Param("id")
	var client models.Client
	if err := database.DB.Where("id = ?", id).First(&client).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := database.DB.Where("id = ? ", id).Delete(&client).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
}

// CreateClient ... Create Client
// @Summary Add a client
// @Description Create new Client based on body parameters
// @Tags Clients
// @Accept json
// @Param Client body CreateAndUpdateClientRequest true "Client Data"
// @Success 200 {object} object
// @Failure 400,401,500 {object} object
// @Router /clients [post]
func CreateClient(c *gin.Context) {
	client := models.Client{}
	if err := c.BindJSON(&client); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := database.DB.Create(&client).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, client)
}
