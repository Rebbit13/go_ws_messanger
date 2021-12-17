package chat

import (
	"go_grpc_messanger/internal/entity"
	"gorm.io/gorm"
)

type RoomController struct {
	db    *gorm.DB
	rooms []Room
}

func (controller *RoomController) flushRooms() {
	var oldRooms []entity.Chat
	controller.db.Find(&oldRooms)
	for _, room := range oldRooms {
		controller.rooms = append(controller.rooms, Room{db: controller.db, ID: room.ID, chat: room})
	}
	return
}

func (controller *RoomController) getRoom(id uint) (room Room, err error) {
	for _, candidateRoom := range controller.rooms {
		if candidateRoom.ID == id {
			room = candidateRoom
			return
		}
	}
	err = &RoomError{"there is no such room"}
	return
}

func (controller *RoomController) CreateNewRoom(title string) (roomEntity entity.Chat, err error) {
	var newChat entity.Chat
	newChat.Title = title
	result := controller.db.Create(&newChat)
	if result.Error != nil {
		err = &RoomError{result.Error.Error()}
		return
	}
	room := Room{db: controller.db, ID: newChat.ID, chat: newChat}
	controller.rooms = append(controller.rooms, room)
	roomEntity = room.chat
	return
}

func (controller *RoomController) GetRoomEntity(id uint) (roomEntity entity.Chat, messages []entity.Message, err error) {
	room, err := controller.getRoom(id)
	if err != nil {
		return
	}
	roomEntity = room.chat
	messages = room.GetMessages()
	return
}

func (controller *RoomController) SendMessage(message entity.Message) (messages []entity.Message, err error) {
	room, err := controller.getRoom(message.ChatID)
	if err != nil {
		return
	}
	err = room.Receive(message)
	if err != nil {
		return
	}
	messages = room.GetMessages()
	return
}

func (controller *RoomController) GetAvailableRooms() (rooms []entity.Chat) {
	for _, room := range controller.rooms {
		rooms = append(rooms, room.chat)
	}
	return
}

func NewRoomController(db *gorm.DB) (controller RoomController) {
	controller = RoomController{db: db}
	controller.flushRooms()
	return
}
