package main

import (
	"fmt"
	"go_grpc_messanger/internal/entity"
	"go_grpc_messanger/internal/storage"
)

func main() {
	entities := []interface{}{&entity.User{}, &entity.Chat{}, &entity.Message{}}
	fmt.Println("create entities")
	var databaseComposer = storage.SqliteDatabase{}
	databaseComposer.InitDatabase(entities)
	fmt.Println("init db")
	db := databaseComposer.GetDatabase()
	fmt.Println("get db")
	fmt.Println(db)
	// Create
	db.Create(&entity.User{Username: "rebbit13", Password: "1234"})

	// Read
	var product entity.User
	db.First(&product, "username=?", "rebbi66t13")
	if product.ID == 0 {
		fmt.Println("There is no such user")
	} else {
		fmt.Println(product)
	}
}
