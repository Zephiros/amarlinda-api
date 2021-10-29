package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Zephiros/amarlinda-api/database"
	"github.com/Zephiros/amarlinda-api/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

type LoginRequest struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

func Register(c *gin.Context) {
	var data map[string]string

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, "")
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	database.DB.Create(&user)

	c.JSON(http.StatusOK, user)
}

// Login ... Login
// @Summary Login
// @Description Login
// @Tags Users
// @Accept json
// @Param login body LoginRequest true "Login Data"
// @Success 200 {object} object
// @Failure 400,401,404 {object} object
// @Router /login [post]
func Login(c *gin.Context) {
	var data map[string]string

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "")
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, &cookie)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

// Logout ... Logout
// @Summary Logout
// @Description Logout
// @Tags Users
// @Success 200 {object} object
// @Failure 400,401,404 {object} object
// @Router /logout [post]
func Logout(c *gin.Context) {
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, &cookie)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
