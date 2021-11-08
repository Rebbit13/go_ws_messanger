package entity

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	Title string
}