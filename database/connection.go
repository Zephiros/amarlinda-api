package database

import (
	"fmt"
	"os"

	"github.com/Zephiros/amarlinda-api/models"
	"github.com/Zephiros/amarlinda-api/pkg/seeds"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	dsn := "" + user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(
		&models.User{},
		&models.Client{},
		&models.Product{},
		&models.Payment{},
		&models.OrderItem{},
		&models.OrderIn{},
		&models.OrderOut{},
	)

	for _, seed := range seeds.All() {
		if err := seed.Run(connection); err != nil {
			fmt.Println("Running seed")
		}
	}
}
