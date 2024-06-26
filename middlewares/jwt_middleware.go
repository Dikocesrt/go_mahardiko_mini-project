package middlewares

import (
	"habit/configs"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userId int) (string, error) {
	//membuat payload
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["role"] = "user"
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	//membuat header
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//verify signature
	return token.SignedString([]byte(configs.InitConfigJWT()))
}

func CreateTokenExpert(userId int) (string, error) {
	//membuat payload
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["role"] = "expert"
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	//membuat header
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//verify signature
	return token.SignedString([]byte(configs.InitConfigJWT()))
}

func CreateTokenAdmin(userId int) (string, error) {
	//membuat payload
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	//membuat header
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//verify signature
	return token.SignedString([]byte(configs.InitConfigJWT()))
}