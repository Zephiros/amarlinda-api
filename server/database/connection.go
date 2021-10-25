package database

import (
    "github.com/Zephiros/amarlinda/models"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    dsn := "admin:admin@tcp(mysql:3306)/amarlinda?charset=utf8mb4&parseTime=True&loc=Local"
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
}
