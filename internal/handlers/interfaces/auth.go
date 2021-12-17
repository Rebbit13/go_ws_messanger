package interfaces

import "go_grpc_messanger/internal/entity"

type Authorization interface {
	SignUp(username, password string) (user entity.User, accessToken, refreshToken string, err error)
	SignIn(username, password string) (user entity.User, accessToken, refreshToken string, err error)
	GetUser(authorizationHeader string) (user entity.User, err error)
	RefreshToken(token string) (accessToken, refreshToken string, err error)
}
