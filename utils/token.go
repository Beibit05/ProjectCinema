package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var JwtKey = []byte("secret")

func GenerateJWT(username string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})
	return token.SignedString(JwtKey)
}
