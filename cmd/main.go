package main

import (
	"go_grpc_messanger/internal/entity"
	"go_grpc_messanger/internal/storage"
)

func main() {
	entities := []interface{}{&entity.User{}, &entity.Chat{}, &entity.Message{}}
	var databaseComposer = storage.SqliteDatabase{}
	databaseComposer.InitDatabase(entities)
	_ = databaseComposer.GetDatabase()
}