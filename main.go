package main

import (
	"fmt"
	"os"

	"github.com/Zephiros/amarlinda-api/database"
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

	database.Connect()

	r := routes.SetupRouter()

	if err := r.Run(":" + os.Getenv("APP_PORT")); err != nil {
		fmt.Println("startup service failed")
	}
}
