package app

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/pat"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

var rd *render.Render

func FacadeHandler() http.Handler {
	mux := pat.New()

	rd = render.New(render.Options{
		Extensions: []string{".html", ".tmpl"},
		Layout:     "hello",
		Directory:  "template",
	})

	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserHandler)

	mux.Get("/hello", helloHandler)

	// 기본
	// mux.Handle("/", http.FileServer(http.Dir("public")))

	n := negroni.Classic()
	n.UseHandler(mux)

	return n
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{FirstName: "Tucker", LastName: "Go", Email: "tuckersgo@gmail.com"}
	rd.JSON(w, http.StatusOK, user)
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		rd.Text(w, http.StatusBadRequest, err.Error())
	}
	user.CreatedAt = time.Now()
	rd.JSON(w, http.StatusOK, user)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	user := User{FirstName: "Tucker", LastName: "Gooooo", Email: "Tuckergoooo@naver.com", CreatedAt: time.Now()}
	rd.HTML(w, http.StatusOK, "body", user)

}
