package middleware

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizationJWT() gin.HandlerFunc {
	return func (c *gin.Context)  {
		cookie, _ := c.Cookie("jwt")

    _, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte("secret"), nil
    })

		if err != nil {
    		c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
