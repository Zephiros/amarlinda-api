package routes

import (
	"github.com/Zephiros/amarlinda-api/controllers"
	"github.com/gin-gonic/gin"
)

func InitClientRoute(r *gin.Engine) {
	clients := r.Group("/clients")
	{
		clients.GET("", controllers.GetClients)
		clients.GET("/:id", controllers.GetClient)
		clients.POST("", controllers.CreateClient)
		clients.PATCH(":id", controllers.UpdateClient)
		clients.DELETE(":id", controllers.DeleteClient)
	}
}
