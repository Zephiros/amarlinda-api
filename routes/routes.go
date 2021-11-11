package routes

import (
	"os"

	"github.com/Zephiros/amarlinda-api/controllers"
	"github.com/Zephiros/amarlinda-api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	if os.Getenv("APP_ENV") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	r.Use(CORSMiddleware())

	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)

	InitSwaggerRoute(r)

	r.Use(middleware.AuthorizationJWT())
	{
		r.POST("/logout", controllers.Logout)
		r.GET("/status", controllers.Status)

		InitUserRoute(r)
		InitProductRoute(r)
		InitClientRoute(r)
		InitPaymentRoute(r)
	}

	return r
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
