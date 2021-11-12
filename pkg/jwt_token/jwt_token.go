package jwt_token

import (
	"github.com/golang-jwt/jwt"
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

func (service *JWTTokenService) validateJWTToken(token string) (claims jwt.Claims, err error) {
	parsedToken, err := jwt.Parse(token, service.validateKey)
	if err != nil {
		return
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !parsedToken.Valid || !ok {
		err = &JWTServiceError{"token invalid"}
		return
	}
	return
}

func (service *JWTTokenService) ValidateAccessToken(token string) (claims jwt.Claims, err error) {
	claims, err = service.validateJWTToken(token)
	if err != nil || claims.(jwt.MapClaims)["type"] != "access" {
		err = &JWTServiceError{"token invalid"}
		return
	}
	return
}

func (service *JWTTokenService) CreateJWTTokens(userId interface{}) (accessToken, refreshToken string) {
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

func (service *JWTTokenService) RefreshToken(token string) (accessToken, refreshToken string, err error) {
	claims, err := service.validateJWTToken(token)
	if err != nil || claims.(jwt.MapClaims)["type"] != "refresh" {
		err = &JWTServiceError{"token invalid"}
		return
	}
	accessToken, refreshToken = service.CreateJWTTokens(claims.(jwt.MapClaims)["user_id"])
	return
}

func NewJWTTokenService(secret []byte, accessTokenTTL time.Duration, refreshTokenTTL time.Duration) (service JWTTokenService, err error){
	if accessTokenTTL >= refreshTokenTTL {
		err = &JWTServiceError{"refresh token must live more than access"}
		return
	}
	return JWTTokenService{secret: secret, accessTokenTTL: accessTokenTTL, refreshTokenTTL: refreshTokenTTL}, nil
}
