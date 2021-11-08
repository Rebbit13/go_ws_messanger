package storage

import "gorm.io/gorm"

type Storage interface {
	InitDatabase()
	GetDatabase() *gorm.DB
}
