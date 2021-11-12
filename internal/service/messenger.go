package service

import "go_grpc_messanger/internal/entity"

type Messenger interface {
	SendMessage(user entity.User, chat entity.Chat, message string) error
}
