package cmd

import (
	"go_grpc_messanger/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqliteDatabase struct {
	db *gorm.DB
}

var dataBase *sqliteDatabase

func InitDataBase() {
	// Init db
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&entity.User{}, &entity.Message{}, &entity.Chat{})
	if err != nil {
		panic("failed to migrate schema")
	}

	dataBase.db = db
}

func GetDatabase() *gorm.DB {
	return dataBase.db
}