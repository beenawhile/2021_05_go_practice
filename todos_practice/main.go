package main

import (
	"net/http"
	"todos_practice/app"

	"github.com/urfave/negroni"
)

func main() {
	mux := app.MakeHandler()

	n := negroni.Classic()

	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
