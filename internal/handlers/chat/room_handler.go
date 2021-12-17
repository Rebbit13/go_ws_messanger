package room

import (
	"github.com/gin-gonic/gin"
	"go_grpc_messanger/internal/entity"
	"go_grpc_messanger/internal/handlers/interfaces"
	"go_grpc_messanger/pkg/json_error_message"
	"go_grpc_messanger/pkg/string_error"
	"strconv"
)

type RoomHandler struct {
	roomService interfaces.Messenger
	authService interfaces.Authorization
}

func (handler *RoomHandler) checkIfAuthorised(context *gin.Context) {
	user, err := handler.authService.GetUser(context.GetHeader("Authorization"))
	if err != nil || user.ID == 0 {
		err = &string_error.StringError{Text: "Token is invalid"}
		context.JSON(401, json_error_message.ErrorMessage{"Token is invalid"})
	}
}

func (handler *RoomHandler) GetAvailableRooms(context *gin.Context) {
	user, err := handler.authService.GetUser(context.GetHeader("Authorization"))
	if err != nil || user.ID == 0 {
		context.JSON(401, json_error_message.ErrorMessage{"Token is invalid"})
	}
	rooms := handler.roomService.GetAvailableRooms()
	context.JSON(200, rooms)
}

func (handler *RoomHandler) CreateNewRoom(context *gin.Context) {
	var room entity.Chat
	err := context.BindJSON(&room)
	if err != nil {
		return
	}
	newRoom, err := handler.roomService.CreateNewRoom(room.Title)
	if err != nil {
		return
	}
	context.JSON(200, newRoom)
}

func (handler *RoomHandler) GetRoom(context *gin.Context) {
	roomIdQuery, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		return
	}
	room, messages, err := handler.roomService.GetRoomEntity(uint(roomIdQuery))
	if err != nil {
		return
	}
	fullRoom := RoomWithMessages{Room: room, Messages: messages}
	context.JSON(200, fullRoom)
}

func (handler *RoomHandler) SendMessage(context *gin.Context) {
	var message entity.Message
	err := context.BindJSON(&message)
	if err != nil {
		return
	}
	messages, err := handler.roomService.SendMessage(message)
	if err != nil {
		return
	}
	context.JSON(200, messages)
}

func BindHandler(roomService interfaces.Messenger, authService interfaces.Authorization, router *gin.Engine) {
	var handler = RoomHandler{roomService: roomService, authService: authService}
	router.GET("/room", handler.GetAvailableRooms)
	router.GET("/room/:id", handler.GetRoom)
	router.POST("/room", handler.CreateNewRoom)
	router.POST("/message", handler.SendMessage)
}
