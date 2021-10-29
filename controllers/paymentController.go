package controllers

import (
	"net/http"

	"github.com/Zephiros/amarlinda-api/database"
	"github.com/Zephiros/amarlinda-api/models"
	"github.com/gin-gonic/gin"
)

// GetPayments ... Get all payments
// @Summary List payments
// @Description Get all payments
// @Tags Payments
// @Success 200 {array} models.Payment
// @Failure 401,404 {object} object
// @Router /payments [get]
func GetPayments(c *gin.Context) {
	payments := []models.Payment{}
	query := database.DB.Select("payments.*").Group("payments.id")
	if err := query.Find(&payments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, payments)
}
