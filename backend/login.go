package main

import (
	"database/sql"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func loginHandler(db *sql.DB, jwtSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {

			err := r.ParseForm()
			if err != nil {
				http.Error(w, "error parsing form", http.StatusBadRequest)
				return
			}

			username := r.FormValue("username")
			password := r.FormValue("password")

			var PasswordHash string
			err = db.QueryRow("SELECT password FROM users WHERE username=$1", username).Scan(&PasswordHash)
			if err != nil {
				http.Error(w, "invalid password or username", http.StatusUnauthorized)
				return
			}

			err = bcrypt.CompareHashAndPassword([]byte(PasswordHash), []byte(password))
			if err != nil {
				http.Error(w, "invalid username or password", http.StatusUnauthorized)
				return
			}

			token, err := CreateToken(username, jwtSecret)

			if err != nil {
				http.Error(w, "failed to create token", http.StatusInternalServerError)
			}
			SetJWTCookie(token, w)
			http.Redirect(w, r, "/profile", http.StatusSeeOther) //maybe need to handle errors
		}
	}
}
