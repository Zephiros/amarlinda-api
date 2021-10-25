package main

import (
    "github.com/Zephiros/amarlinda/database"
    "github.com/Zephiros/amarlinda/routes"
    "fmt"
)

func main() {
    database.Connect()

    r := routes.SetupRouter()

    if err := r.Run(); err != nil {
  		  fmt.Println("startup service failed, err:%v\n", err)
  	}
}
