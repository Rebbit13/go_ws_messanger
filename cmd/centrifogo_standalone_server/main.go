package main

import (
	"github.com/gin-gonic/gin"
	"go_grpc_messanger/internal/entity"
	"go_grpc_messanger/internal/handlers/auth"
	"go_grpc_messanger/internal/handlers/centrifuge_messenger"
	room "go_grpc_messanger/internal/handlers/chat"
	"go_grpc_messanger/internal/handlers/page"
	"go_grpc_messanger/internal/service/authorization"
	"go_grpc_messanger/internal/service/chat"
	"go_grpc_messanger/internal/storage"
	"os"
	"time"
)

func main() {
	entities := []interface{}{&entity.User{}, &entity.Chat{}, &entity.Message{}}
	var databaseComposer = storage.SqliteDatabase{}
	databaseComposer.InitDatabase(entities)
	db := databaseComposer.GetDatabase()
	authService, err := authorization.NewJWTAuth(db, []byte(os.Getenv("JWT_SECRET")), time.Duration(20*time.Minute), time.Duration(20*time.Minute))
	roomService := chat.NewRoomController(db)
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	auth.BindHandler(&authService, r)
	page.BindHandler(r)
	room.BindHandler(&roomService, &authService, r)
	centrifuge_messenger.BindHandler(&roomService, &authService, r)
	r.Run()
}
