package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// User Struct
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// user 정보를 담고 있는 map
var userMap map[int]*User
var lastID int

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello, world")
}
func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get UserInfo by /users/{id}")
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	// mux.Vars 안에 request 넣어주면 알아서 파싱해줌
	// vars := mux.Vars(r)
	// fmt.Fprint(w, "User Id:", vars["id"])

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	user, ok := userMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User Id:", id)
		return
	}
	// user := new(User)

	// user.ID = lastID
	// user.FirstName = "tucker"
	// user.LastName = "kim"
	// user.Email = "tucker@naver.com"

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	// create user
	lastID++
	user.ID = lastID
	user.CreatedAt = time.Now()

	userMap[user.ID] = user

	w.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

// NewHandler make a new handler
func NewHandler() http.Handler {

	// user초기화
	userMap = make(map[int]*User)
	lastID = 0

	mux := mux.NewRouter()
	// mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", userHandler).Methods("GET")
	mux.HandleFunc("/users", createUserHandler).Methods("POST")
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler)

	// 따로 지정하지 않으면 다른 uri 가 들어오더라도 테스트 통과하게 됨
	// 예외 처리

	return mux
}
