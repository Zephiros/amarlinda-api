package main

import (
	"fmt"
	"os"

	"github.com/Zephiros/amarlinda-api/database"
	"github.com/Zephiros/amarlinda-api/docs"
	"github.com/Zephiros/amarlinda-api/routes"
	"github.com/joho/godotenv"
)

// @title API documentation
// @version 1.0.0

// @host localhost:8082
// @BasePath /api
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	docs.SwaggerInfo.Title = "Amarlinda Store API"
	docs.SwaggerInfo.Description = "This is a Amarlinda API documentation."
	docs.SwaggerInfo.Version = os.Getenv("APP_VERSION")
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	database.Connect()

	r := routes.SetupRouter()

	if err := r.Run(":" + os.Getenv("APP_PORT")); err != nil {
		fmt.Println("startup service failed")
	}
}
