package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_grpc_messanger/internal/entity"
	"go_grpc_messanger/internal/handlers/interfaces"
)

type AuthHandler struct {
	authService interfaces.Authorization
}

func (handler *AuthHandler) SignUp(context *gin.Context) {
	var user = entity.User{}
	err := context.BindJSON(&user)
	if err != nil {
		return
	}
	user, access, refresh, err := handler.authService.SignUp(user.Username, user.Password)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	context.JSON(200, TokenPair{access, refresh})
}

func (handler *AuthHandler) SignIn(context *gin.Context) {
	var user = entity.User{}
	err := context.BindJSON(&user)
	if err != nil {
		return
	}
	user, access, refresh, err := handler.authService.SignIn(user.Username, user.Password)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	context.JSON(200, TokenPair{access, refresh})
}

func (handler *AuthHandler) Refresh(context *gin.Context) {
	var refresh = RefreshToken{}
	err := context.BindJSON(&refresh)
	if err != nil {
		return
	}
	accessToken, refreshToken, err := handler.authService.RefreshToken(refresh.RefreshToken)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	context.JSON(200, TokenPair{accessToken, refreshToken})
}

func BindHandler(authService interfaces.Authorization, router *gin.Engine) {
	var handler = AuthHandler{authService: authService}
	router.POST("/sign_in", handler.SignIn)
	router.POST("/sign_up", handler.SignUp)
	router.POST("/refresh", handler.Refresh)
}
