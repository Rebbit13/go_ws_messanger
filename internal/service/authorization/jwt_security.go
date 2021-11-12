package authorization

import (
	"github.com/golang-jwt/jwt"
	"go_grpc_messanger/internal/entity"
	"go_grpc_messanger/pkg/jwt_token"
	"gorm.io/gorm"
	"time"
)

type JWTAuth struct {
	db *gorm.DB
	jwtService jwt_token.JWTTokenService
}

func NewJWTAuth(db *gorm.DB, secret []byte, accessTokenTTL time.Duration, refreshTokenTTL time.Duration) (auth JWTAuth, err error){
	jwtService, err := jwt_token.NewJWTTokenService(secret, accessTokenTTL, refreshTokenTTL)
	if err != nil {
		return
	}
	return JWTAuth{db: db, jwtService: jwtService}, nil
}

func (auth *JWTAuth) getUser(token string) (user entity.User, err error) {
	claims, err := auth.jwtService.ValidateJWTToken(token)
	auth.db.First(&user, "id = ?", claims.(jwt.MapClaims)["user_id"])
	return
}

func (auth *JWTAuth) SignIn()  {
	
}

func (auth *JWTAuth) SignUp(username string, password string) (err error, accessToken string, refreshToken string) {
	return
}
