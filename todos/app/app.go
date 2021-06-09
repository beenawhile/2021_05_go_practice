package app

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rd *render.Render

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

var todoMap map[int]*Todo

func MakeHandler() http.Handler {

	todoMap = make(map[int]*Todo)

	addTestTodos()

	rd = render.New()

	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/todos", getTodoListHandler).Methods("GET")

	return r
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := []*Todo{}

	for _, v := range todoMap {
		list = append(list, v)
	}
	rd.JSON(w, http.StatusOK, list)
}

// 테스트용 데이터
func addTestTodos() {
	todoMap[1] = &Todo{ID: 1, Name: "Buy a milk", Completed: false, CreatedAt: time.Now()}
	todoMap[2] = &Todo{ID: 2, Name: "Exercise", Completed: true, CreatedAt: time.Now()}
	todoMap[3] = &Todo{ID: 3, Name: "Home work", Completed: false, CreatedAt: time.Now()}
}
