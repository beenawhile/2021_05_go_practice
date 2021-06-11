package app

import (
	"net/http"
	"strconv"
	"todos/model"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rd *render.Render = render.New()

type AppHandler struct {
	// interface를 포함하고 있다 => 상속하고 다른 의미 (has 관계이지 is 관계가 아님)
	http.Handler
	db model.DBHandler
}

func MakeHandler(filepath string) *AppHandler {

	r := mux.NewRouter()
	a := &AppHandler{
		Handler: r,
		db:      model.NewDBHandler(filepath),
	}

	r.HandleFunc("/", a.indexHandler)
	r.HandleFunc("/todos", a.getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos", a.addTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", a.removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/complete-todo/{id:[0-9]+}", a.completeTodoHandler).Methods("GET")

	return a
}

func (a *AppHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func (a *AppHandler) getTodoListHandler(w http.ResponseWriter, r *http.Request) {

	list := a.db.GetTodos()
	rd.JSON(w, http.StatusOK, list)
}

func (a *AppHandler) addTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	todo := a.db.AddTodo(name)

	rd.JSON(w, http.StatusCreated, todo)
}

func (a *AppHandler) removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	// id 값 받아오기
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	ok := a.db.RemoveTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{Success: true})
	} else {
		rd.JSON(w, http.StatusOK, Success{Success: false})
	}
}

// 응답결과 받는 structure
type Success struct {
	Success bool `json:"success"`
}

func (a *AppHandler) completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	// id 값 받기
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("complete") == "true"

	ok := a.db.CompleteTodo(id, complete)
	if ok {
		rd.JSON(w, http.StatusOK, Success{Success: true})
	} else {
		rd.JSON(w, http.StatusOK, Success{Success: false})
	}
}

func (a *AppHandler) Close() {
	a.db.Close()
}
