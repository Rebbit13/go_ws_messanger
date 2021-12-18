package chat

import (
	"go_grpc_messanger/internal/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Room struct {
	db      *gorm.DB
	ID      uint
	chat    entity.Chat
	channel []chan entity.Message
}

func (chat *Room) GetMessages() (messages []entity.Message) {
	chat.db.Preload(clause.Associations).Where("chat_id = ?", chat.ID).Find(&messages)
	return
}

func (chat *Room) Receive(message entity.Message) error {
	chat.db.Create(&message)
	return nil
}
