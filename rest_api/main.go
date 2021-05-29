package main

import (
	"net/http"
	"rest_api/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
