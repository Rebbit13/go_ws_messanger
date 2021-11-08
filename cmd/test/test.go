package main

import (
	"fmt"
	"go_grpc_messanger/internal/entity"
	"go_grpc_messanger/pkg/database"
)

func main() {
	database.InitDataBase()
	db := database.GetDatabase()
	// Create
	db.Create(&entity.User{Username: "rebbit13", Password: "1234"})

	// Read
	var product entity.User
	db.First(&product, "username=?", "rebbi66t13")
	if product.ID == 0 {

	}
	fmt.Println(product)
}
