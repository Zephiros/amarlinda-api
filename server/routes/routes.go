package routes

import (
    "github.com/Zephiros/amarlinda/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.POST("/api/register", controllers.Register)
    r.POST("/api/login", controllers.Login)
    r.POST("/api/logout", controllers.Logout)
    r.GET("/api/user", controllers.User)

    return r
}
