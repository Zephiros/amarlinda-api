package controllers

import (
    "github.com/Zephiros/amarlinda/database"
    "github.com/Zephiros/amarlinda/models"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
    "strconv"
    "time"
    "net/http"
)

const SecretKey = "secret"

func Register(c *gin.Context) {
    var data map[string]string

    if err := c.ShouldBindJSON(&data); err != nil {
       c.JSON(http.StatusBadRequest, "")
    }

    password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

    user := models.User{
      Name: data["name"],
      Email: data["email"],
      Password: password,
    }

    database.DB.Create(&user)

    c.JSON(http.StatusOK, user);
}

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
        Issuer: strconv.Itoa(int(user.Id)),
        ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
    })

    token, err := claims.SignedString([]byte(SecretKey))

    if err != nil {
        c.JSON(http.StatusInternalServerError, nil)
        return
    }

    cookie := http.Cookie{
        Name: "jwt",
        Value: token,
        Expires: time.Now().Add(time.Hour),
        HttpOnly: true,
    }

    http.SetCookie(c.Writer, &cookie)

    c.JSON(http.StatusOK, gin.H{
        "message": "success",
    })
}

func Logout(c *gin.Context) {
    cookie := http.Cookie{
        Name: "jwt",
        Value: "",
        Expires: time.Now().Add(-time.Hour),
        HttpOnly: true,
    }

    http.SetCookie(c.Writer, &cookie)

    c.JSON(http.StatusOK, gin.H{
        "message": "success",
    })
}

func User(c *gin.Context) {
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
