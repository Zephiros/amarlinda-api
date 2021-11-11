package routes

import (
	"github.com/Zephiros/amarlinda-api/controllers"
	"github.com/gin-gonic/gin"
)

func InitPaymentRoute(r *gin.Engine) {
	payments := r.Group("/payments")
	{
		payments.GET("", controllers.GetPayments)
	}
}
