package main

import (
	"net/http"
	"web/myapp"
)

func main() {

	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
