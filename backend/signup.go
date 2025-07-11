package main

import (
	"database/sql"
	"net/http"
)

func signupHandler(db *sql.DB, jwtSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing form", http.StatusBadRequest)
				return
			}

			name := r.FormValue("name")
			username := r.FormValue("username")

			//hashing password
			hashedPassword, err := hashPassword(r.FormValue("password"))

			if err != nil {
				http.Error(w, "failed to hash password", http.StatusInternalServerError)
			}

			// line := fmt.Sprintf("Name: %s, Username: %s, Password: %s\n", name, username, hashedPassword)
			//now we will append to a file for testing
			// file, err := os.OpenFile("./testdata.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //I copied the  last 2 argument, will check later
			// if err != nil {
			// 	http.Error(w, "error while opening file", http.StatusInternalServerError)
			// }
			// defer file.Close()
			// _, err = file.WriteString(line)

			// if err != nil {
			// 	http.Error(w, "error while saving credentials", http.StatusInternalServerError)
			// 	return
			// }

			//inserting form data to database
			_, err = db.Exec(`INSERT INTO users (name, username, password) VALUES ($1, $2, $3)`, name, username, hashedPassword)
			if err != nil {
				http.Error(w, "DB insert failed", http.StatusInternalServerError)
				return
			}

			//creating jwt
			token, err := CreateToken(username, jwtSecret)
			if err != nil {
				http.Error(w, "failed to generate jwt", http.StatusInternalServerError)
			}
			SetJWTCookie(token, w)
			http.Redirect(w, r, "/profile", http.StatusSeeOther)
		}
	}
}
