package main

import (
	"net/http"
	"render_pat_practice/app"
	"time"
)

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	// mux := http.NewServeMux()

	// rest api 1
	// mux := mux.NewRouter()

	http.ListenAndServe(":3000", app.FacadeHandler())
}
