package interfaces

import "go_grpc_messanger/internal/entity"

type Messenger interface {
	GetAvailableRooms() (rooms []entity.Chat)
	NewRoom(title string) (roomEntity entity.Chat, err error)
	GetRoomEntity(id uint) (roomEntity entity.Chat, messages []entity.Message, err error)
	SendMessage(message entity.Message) (messages []entity.Message, err error)
}
