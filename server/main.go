package main

import (
	"fmt"

	"github.com/Zephiros/amarlinda/database"
	"github.com/Zephiros/amarlinda/docs"
	"github.com/Zephiros/amarlinda/routes"
)

// @title API documentation
// @version 1.0.0

// @host localhost:8082
// @BasePath /api
func main() {
	docs.SwaggerInfo.Title = "Amarlinda Store Server API"
	docs.SwaggerInfo.Description = "This is a Amarlinda API documentation."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	database.Connect()

	r := routes.SetupRouter()

	if err := r.Run(":8082"); err != nil {
		fmt.Println("startup service failed")
	}
}
