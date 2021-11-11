package routes

import (
	"os"

	"github.com/Zephiros/amarlinda-api/docs"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitSwaggerRoute(r *gin.Engine) {
	docs.SwaggerInfo.Title = "Amarlinda Store API"
	docs.SwaggerInfo.Description = "This is a Amarlinda API documentation."
	docs.SwaggerInfo.Version = os.Getenv("APP_VERSION")
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
