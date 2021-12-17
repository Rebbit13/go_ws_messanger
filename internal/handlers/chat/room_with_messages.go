package room

import "go_grpc_messanger/internal/entity"

type RoomWithMessages struct {
	Room     entity.Chat
	Messages []entity.Message
}
