package main

import (
	"net/http"
	"todos_practice/app"

	"github.com/urfave/negroni"
)

func main() {

	m := app.MakeHandler("./test.db")
	defer m.Close()

	n := negroni.Classic()

	n.UseHandler(m)

	http.ListenAndServe(":3000", n)

}
