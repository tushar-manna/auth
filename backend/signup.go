package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			log.Println(err)
			return
		}

		name := r.FormValue("name")
		username := r.FormValue("username")
		password := r.FormValue("password")
		line := fmt.Sprintf("Name: %s, Username: %s, Password: %s\n", name, username, password)

		//now we will append to a file for testing
		file, err := os.OpenFile("./testdata.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //I copied the  last 2 argument, will check later
		if err != nil {
			http.Error(w, "error while opening file", http.StatusInternalServerError)
			log.Println(err)
		}

		defer file.Close()

		test, err := file.WriteString(line)

		if err != nil {
			http.Error(w, "error while saving credentials", http.StatusInternalServerError)
			log.Println(test)
			log.Println(err)
			return
		}
		fmt.Fprintln(w, "signup saved!")
	}
}
