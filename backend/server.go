package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", fs)
	http.HandleFunc("/signup", signupHandler)
	http.ListenAndServe(":8000", nil)
	fmt.Println("Server is running on port 8000")
}
