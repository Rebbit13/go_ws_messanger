package storage

import "gorm.io/gorm"

type Storage interface {
	InitDataBase()
	GetDatabase() *gorm.DB
}
