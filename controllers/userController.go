package controllers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Zephiros/amarlinda-api/database"
	"github.com/Zephiros/amarlinda-api/helpers"
	"github.com/Zephiros/amarlinda-api/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UpdateUserRequest struct {
	Name           string `json:"name" binding:"required"`
	Responsability string `json:"responsability" binding:"required"`
}

type UpdateUserPasswordRequest struct {
	Password string `json:"password" binding:"required"`
	Confirm  string `json:"confirm" binding:"required"`
}

func GetUser(c *gin.Context) (models.User, error) {
	cookie, _ := c.Cookie("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return models.User{}, err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return user, nil
}

// GetUserProfile ... Get User data
// @Summary Get user
// @Description Get logged user data
// @Tags Users
// @Success 200 {object} models.User
// @Failure 400,401,404 {object} object
// @Router /users/profile [get]
func GetUserProfile(c *gin.Context) {
	user, err := GetUser(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, nil)
	}

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
	user, err := GetUser(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, nil)
	}

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

// UpdateUserPassword ... Update User password
// @Summary Update password
// @Description Update logged user password
// @Tags Users
// @Param User body UpdateUserPasswordRequest true "User password"
// @Success 200 {object} models.User
// @Failure 400,401,404 {object} object
// @Router /users/password [patch]
func UpdateUserPassword(c *gin.Context) {
	user, err := GetUser(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, nil)
	}

	var data UpdateUserPasswordRequest

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if data.Password != data.Confirm {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	if !helpers.Validate(data.Password) {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	user.Password = password

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUserAvatar ... Update User avatar
// @Summary Update avatar
// @Description Update logged user avatar
// @Tags Users
// @Param file formData file true "User avatar"
// @Success 200 {object} models.User
// @Failure 400,401,404 {object} object
// @Router /users/avatar [patch]
func UpdateUserAvatar(c *gin.Context) {
	user, err := GetUser(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, nil)
	}

	file, header, err := c.Request.FormFile("file")
	filename := header.Filename

	dirPath := fmt.Sprintf("uploads/%d", user.Id)

	if _, err := os.Stat(dirPath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(dirPath, os.ModePerm)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	saveFilePath := fmt.Sprintf("%s/%s", dirPath, filename)

	out, err := os.Create(saveFilePath)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	defer out.Close()
	_, err = io.Copy(out, file)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	user.Avatar = filename

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

// @TODO GET LOGGED USER
