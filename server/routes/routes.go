package routes

import (
	"github.com/Zephiros/amarlinda/controllers"
	"github.com/Zephiros/amarlinda/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(CORSMiddleware())

	r.POST("/api/register", controllers.Register)
	r.POST("/api/login", controllers.Login)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(middleware.AuthorizationJWT())
	{
		r.POST("/api/logout", controllers.Logout)
		r.GET("/api/user", controllers.User)
		r.GET("/api/status", controllers.Status)

		products := r.Group("/api/products")
		{
			products.GET("", controllers.GetProducts)
			products.GET("/:id", controllers.GetProduct)
			products.POST("", controllers.CreateProduct)
			products.PATCH(":id", controllers.UpdateProduct)
			products.DELETE(":id", controllers.DeleteProduct)
		}
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
