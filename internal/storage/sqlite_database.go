package storage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteDatabase struct {
	db *gorm.DB
}

func (database *SqliteDatabase) InitDatabase(entities []interface{}) {
	// Init db
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	for _, model := range entities {
		err = db.AutoMigrate(model)
		if err != nil {
			panic("failed to migrate schema")
		}
	}
	database.db = db
}

func (database *SqliteDatabase) GetDatabase() *gorm.DB {
	return database.db
}
