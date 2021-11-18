package routes

import (
	"github.com/Zephiros/amarlinda-api/controllers"
	"github.com/gin-gonic/gin"
)

func InitProductRoute(r *gin.Engine) {
	products := r.Group("/products")
	{
		products.GET("", controllers.GetProducts)
		products.GET("/:id", controllers.GetProduct)
		products.POST("", controllers.CreateProduct)
		products.PATCH(":id", controllers.UpdateProduct)
		products.DELETE(":id", controllers.DeleteProduct)
		products.POST("/import", controllers.ImportProduct)
	}
}
