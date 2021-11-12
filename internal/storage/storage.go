package storage

import "gorm.io/gorm"

type Storage interface {
	InitDatabase(entities []interface{})
	GetDatabase() *gorm.DB
}
