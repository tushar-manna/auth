package main

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(bytes), err
}

func SetJWTCookie(JWT string, w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    JWT,
		Path:     "/",
		MaxAge:   24 * 60 * 60 * 7,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
}
