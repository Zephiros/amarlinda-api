package routes

import (
	"os"

	"github.com/Zephiros/amarlinda-api/controllers"
	"github.com/Zephiros/amarlinda-api/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetupRouter() *gin.Engine {
	if os.Getenv("APP_ENV") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	r.Use(CORSMiddleware())

	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(middleware.AuthorizationJWT())
	{
		r.POST("/logout", controllers.Logout)
		r.GET("/status", controllers.Status)

		users := r.Group("/users")
		{
			users.GET("/profile", controllers.GetUserProfile)
			users.PATCH("/profile", controllers.UpdateUserProfile)
			users.PATCH("/password", controllers.UpdateUserPassword)
			users.PATCH("/avatar", controllers.UpdateUserAvatar)
		}

		products := r.Group("/products")
		{
			products.GET("", controllers.GetProducts)
			products.GET("/:id", controllers.GetProduct)
			products.POST("", controllers.CreateProduct)
			products.PATCH(":id", controllers.UpdateProduct)
			products.DELETE(":id", controllers.DeleteProduct)
		}

		clients := r.Group("/clients")
		{
			clients.GET("", controllers.GetClients)
			clients.GET("/:id", controllers.GetClient)
			clients.POST("", controllers.CreateClient)
			clients.PATCH(":id", controllers.UpdateClient)
			clients.DELETE(":id", controllers.DeleteClient)
		}

		payments := r.Group("/payments")
		{
			payments.GET("", controllers.GetPayments)
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
