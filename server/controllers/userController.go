package controllers

import (
	"net/http"

	"github.com/Zephiros/amarlinda/database"
	"github.com/Zephiros/amarlinda/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UpdateUserRequest struct {
	Name           string `json:"name" binding:"required"`
	Responsability string `json:"responsability" binding:"required"`
}

// GetUserProfile ... Get User data
// @Summary Get user
// @Description Get logged user data
// @Tags Users
// @Success 200 {object} models.User
// @Failure 400,401,404 {object} object
// @Router /users/profile [get]
func GetUserProfile(c *gin.Context) {
	cookie, _ := c.Cookie("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	c.JSON(http.StatusOK, user)
}

// UpdateUserProfile ... Update User data
// @Summary Update user
// @Description Update logged user data
// @Tags Users
// @Param User body UpdateUserRequest true "User Data"
// @Success 200 {object} models.User
// @Failure 400,401,404 {object} object
// @Router /users/profile [patch]
func UpdateUserProfile(c *gin.Context) {
	cookie, _ := c.Cookie("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
