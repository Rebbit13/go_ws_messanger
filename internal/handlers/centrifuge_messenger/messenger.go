package centrifuge_messenger

import (
	"encoding/json"
	"fmt"
	"github.com/centrifugal/centrifuge"
	"github.com/gin-gonic/gin"
	"go_grpc_messanger/internal/entity"
	"go_grpc_messanger/internal/handlers/interfaces"
	"go_grpc_messanger/pkg/json_error_message"
	"go_grpc_messanger/pkg/string_error"
	"net/http"
	"strconv"
)

type CentrifugeBroker struct {
	node        *centrifuge.Node
	authService interfaces.Authorization
	roomService interfaces.Messenger
}

func (broker *CentrifugeBroker) auth(h http.Handler) gin.HandlerFunc {
	return func(context *gin.Context) {
		auth, err := context.Cookie("access_token")
		user, err := broker.authService.GetUser(fmt.Sprintf("Bearer %s", auth))
		if err != nil || user.ID == 0 {
			err = &string_error.StringError{Text: "Token is invalid"}
			context.JSON(401, json_error_message.ErrorMessage{"Token is invalid"})
		}
		userPacked, _ := json.Marshal(user)
		cred := &centrifuge.Credentials{
			UserID: user.Username,
			Info:   userPacked,
		}
		newCtx := centrifuge.SetCredentials(context, cred)
		request := context.Request.WithContext(newCtx)
		h.ServeHTTP(context.Writer, request)
	}
}

func (broker *CentrifugeBroker) GetWSHandler() gin.HandlerFunc {
	wsHandler := centrifuge.NewWebsocketHandler(broker.node, centrifuge.WebsocketConfig{})
	return broker.auth(wsHandler)
}

func (broker *CentrifugeBroker) Start() (err error) {
	err = broker.node.Run()
	return
}

func (broker *CentrifugeBroker) initNodeConfiguration() {
	broker.node.OnConnecting(onConnecting)

	broker.node.OnConnect(func(client *centrifuge.Client) {

		client.OnSubscribe(func(e centrifuge.SubscribeEvent, cb centrifuge.SubscribeCallback) {
			cb(centrifuge.SubscribeReply{}, nil)
			roomIdQuery, _ := strconv.ParseUint(e.Channel, 10, 64)
			_, messages, _ := broker.roomService.GetRoomEntity(uint(roomIdQuery))
			for _, oldMessage := range messages {
				data, _ := json.Marshal(oldMessage)
				_, _ = broker.node.Publish(
					e.Channel, data,
				)
			}
		})

		client.OnPublish(func(e centrifuge.PublishEvent, cb centrifuge.PublishCallback) {
			var user entity.User
			err := json.Unmarshal(client.Info(), &user)
			roomIdQuery, _ := strconv.ParseUint(e.Channel, 10, 64)
			if err != nil {
				fmt.Println(err)
			}
			messages, _ := broker.roomService.SendMessage(
				user.ID, uint(roomIdQuery), string(e.Data))
			data, _ := json.Marshal(messages[len(messages)-1])
			result, _ := broker.node.Publish(
				e.Channel, data,
			)
			cb(centrifuge.PublishReply{Result: &result}, err)
		})

	})
}

func BindHandler(roomService interfaces.Messenger, authService interfaces.Authorization, router *gin.Engine) CentrifugeBroker {
	cfg := centrifuge.DefaultConfig
	node, err := centrifuge.New(cfg)
	if err != nil {
		panic(err)
	}
	broker := CentrifugeBroker{node: node, authService: authService, roomService: roomService}
	broker.initNodeConfiguration()
	router.GET("/connection/websocket", broker.GetWSHandler())
	err = broker.Start()
	if err != nil {
		panic(err)
	}
	return broker
}
