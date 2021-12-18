package interfaces

import "go_grpc_messanger/internal/entity"

type Messenger interface {
	GetAvailableRooms() (rooms []entity.Chat)
	CreateNewRoom(title string) (roomEntity entity.Chat, err error)
	GetRoomEntity(id uint) (roomEntity entity.Chat, messages []entity.Message, err error)
	SendMessage(userID uint, chatID uint, text string) (messages []entity.Message, err error)
}
