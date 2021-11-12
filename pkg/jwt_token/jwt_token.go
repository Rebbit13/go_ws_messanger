package jwt_token

import (
	"github.com/golang-jwt/jwt"
	"go_grpc_messanger/pkg/string_error"
	"time"
)

type JWTTokenService struct {
	secret []byte
	accessTokenTTL time.Duration
	refreshTokenTTL time.Duration
}

func (service *JWTTokenService) validateKey(*jwt.Token) (interface{}, error) {
	return service.secret, nil
}

func (service *JWTTokenService) ValidateJWTToken(token string) (claims jwt.Claims, err error) {
	parsedToken, err := jwt.Parse(token, service.validateKey)
	if err != nil {
		return
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !parsedToken.Valid || !ok {
		err = &string_error.StringError{"token invalid"}
		return
	}
	return
}

func (service *JWTTokenService) CreateJWTTokens(userId string) (accessToken, refreshToken string) {
	accessToken, _ = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"type":     "access",
		"user_id":  userId,
		"ttl":      time.Now().Add(service.accessTokenTTL),
	}).SignedString(service.secret)
	refreshToken, _ = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"type":     "refresh",
		"user_id":  userId,
		"ttl":      time.Now().Add(service.refreshTokenTTL),
	}).SignedString(service.secret)
	return
}

func NewJWTTokenService(secret []byte, accessTokenTTL time.Duration, refreshTokenTTL time.Duration) (service JWTTokenService, err error){
	if accessTokenTTL >= refreshTokenTTL {
		err = &string_error.StringError{"refresh token must live more than access"}
		return
	}
	return JWTTokenService{secret: secret, accessTokenTTL: accessTokenTTL, refreshTokenTTL: refreshTokenTTL}, nil
}
