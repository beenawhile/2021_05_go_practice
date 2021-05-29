package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type fooHandler struct{}

// json 담을 struct
type User struct {
	// go convention 이니까 참고
	// ``로 annotation 붙여줌
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt time.Time
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request: ", err)
		return
	}

	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user)
	// header에 어떤 형식인지 알려주기
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(data))
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "World"
	}

	fmt.Fprintf(w, "Hello %s!", name)
}

func NewHttpHandler() http.Handler {
	// Mux == Router(in go)
	// 라우터를 만들어서 동적으로 전달하는 방법으로 해보자
	mux := http.NewServeMux()

	// handlefunc 는 함수 직점 넘김
	mux.HandleFunc("/", indexHandler)

	mux.HandleFunc("/bar", barHandler)

	// handle 은 인스턴스 형태로 넘김
	mux.Handle("/foo", &fooHandler{})

	return mux
}
