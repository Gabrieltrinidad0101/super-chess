package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type PlayerJwt struct {
	PlayerId uint
	jwt.StandardClaims
}

func GetJwtSecret() []byte {
	conf := Enviroments()
	return []byte(conf.JwtSecret)
}

type JsonWebToken struct{}

func (JsonWebToken) CreateToken(playerId uint) (string, error) {
	userJwt := &PlayerJwt{
		PlayerId: playerId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userJwt)
	tokenString, err := token.SignedString(GetJwtSecret())

	return tokenString, err
}
