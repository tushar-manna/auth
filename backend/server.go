package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env")
	}

	JWT_SECRET := os.Getenv("JWT_SECRET")
	if JWT_SECRET == "" {
		log.Fatal("failed to load JWT_SECRET")
	}

	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", fs)
	http.HandleFunc("/signup", signupHandler(db, JWT_SECRET))
	http.HandleFunc("/login", loginHandler(db, JWT_SECRET))
	http.Handle("/profile", AuthMiddleWare(http.HandlerFunc(ProfileHandler), JWT_SECRET))
	http.ListenAndServe(":8000", nil)
	fmt.Println("Server is running on port 8000")
}
