package authorization

import (
	"github.com/golang-jwt/jwt"
	"go_grpc_messanger/internal/entity"
	"go_grpc_messanger/pkg/jwt_token"
	"go_grpc_messanger/pkg/password_hash"
	"gorm.io/gorm"
	"time"
)

type JWTAuth struct {
	db          *gorm.DB
	jwtService  jwt_token.JWTTokenService
	hashService password_hash.HashService
}

func (auth *JWTAuth) getUser(token string) (user entity.User, err error) {
	claims, err := auth.jwtService.ValidateAccessToken(token)
	if err != nil {
		return
	}
	auth.db.First(&user, "id = ?", claims.(jwt.MapClaims)["user_id"])
	if user.ID == 0 {
		err = &AuthError{"user does not exist"}
		return
	}
	return
}

func (auth *JWTAuth) SignIn(username, password string) (user entity.User, accessToken, refreshToken string, err error) {
	auth.db.First(&user, "username = ?", username)
	if user.ID == 0 {
		err = &AuthError{"user does not exist"}
		return
	}
	isPasswordValid := auth.hashService.CheckPasswordHash(password, user.Password)
	if !isPasswordValid {
		err = &AuthError{"password invalid"}
		return
	}
	accessToken, refreshToken = auth.jwtService.CreateJWTTokens(user.ID)
	return
}

func (auth *JWTAuth) SignUp(username, password string) (user entity.User, accessToken, refreshToken string, err error) {
	auth.db.First(&user, "username = ?", username)
	if user.ID != 0 {
		err = &AuthError{"user already exist"}
		return
	}
	hashedPassword, err := auth.hashService.HashPassword(password)
	if err != nil {
		err = &AuthError{err.Error()}
		return
	}
	user = entity.User{Username: username, Password: hashedPassword}
	result := auth.db.Create(&user)
	if result.Error != nil {
		err = &AuthError{result.Error.Error()}
		return
	}
	if user.ID == 0 {
		err = &AuthError{"cannot create user"}
		return
	}
	accessToken, refreshToken = auth.jwtService.CreateJWTTokens(user.ID)
	return
}

func (auth *JWTAuth) GetUser(authorizationHeader string) (user entity.User, err error) {
	token := auth.jwtService.GetTokenFromHeader(authorizationHeader)
	return auth.getUser(token)
}

func (auth *JWTAuth) RefreshToken(token string) (accessToken, refreshToken string, err error) {
	return auth.jwtService.RefreshToken(token)
}

func NewJWTAuth(db *gorm.DB, secret []byte, accessTokenTTL time.Duration, refreshTokenTTL time.Duration) (auth JWTAuth, err error) {
	jwtService, err := jwt_token.NewJWTTokenService(secret, accessTokenTTL, refreshTokenTTL)
	hashService := password_hash.NewHashService(14)
	if err != nil {
		return
	}
	return JWTAuth{db: db, jwtService: jwtService, hashService: hashService}, nil
}
