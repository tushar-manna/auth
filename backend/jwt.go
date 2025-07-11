package main

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(username, secret string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS384, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), //Copy pasting, I will look at the time module later
	})

	return token.SignedString([]byte(secret))
}
