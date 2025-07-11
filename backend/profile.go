package main

import (
	"fmt"
	"net/http"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)
	fmt.Fprintf(w, "hey %s! yes it's a protected route and you can only view it if your are logged in", username)
}
