package main

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleWare(next http.Handler, JWTSecret string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil || cookie.Value == "" {
			http.Error(w, "unauthorized: no auth token found", http.StatusUnauthorized)
			return
		}

		tokenStr := cookie.Value

		type MyClaims struct {
			Username string `json:"username"` //need to see this pattern later
			jwt.RegisteredClaims
		}

		claims := &MyClaims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {

			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(JWTSecret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "username", claims.Username)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
