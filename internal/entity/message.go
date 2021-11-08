package entity

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	gorm.Model
	CreatedAt time.Time
	UserID uint
	User   User
	ChatID uint
	Chat   Chat
	Text   string
}
