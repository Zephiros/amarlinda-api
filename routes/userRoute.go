package routes

import (
	"github.com/Zephiros/amarlinda-api/controllers"
	"github.com/gin-gonic/gin"
)

func InitUserRoute(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("/profile", controllers.GetUserProfile)
		users.PATCH("/profile", controllers.UpdateUserProfile)
		users.PATCH("/password", controllers.UpdateUserPassword)
		users.PATCH("/avatar", controllers.UpdateUserAvatar)
	}
}
