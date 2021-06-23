package app

import (
	"net/http"
	"strconv"
	"todos_practice/model"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

type AppHandler struct {
	http.Handler
	db model.DBHandler
}

type Success struct {
	IsSuccess bool `json:"isSuccess"`
}

var rd *render.Render = render.New()

func MakeHandler() AppHandler {

	m := mux.NewRouter()

	a := &AppHandler{
		Handler: m,
		db:      model.NewDBHandler(),
	}

	m.HandleFunc("/", a.indexHandler)
	m.HandleFunc("/todos", a.getTodoListHandler).Methods("GET")
	m.HandleFunc("/todos", a.addTodoHandler).Methods("POST")
	m.HandleFunc("/todos/{id:[0-9]+}", a.removeTodoHandler).Methods("DELETE")
	m.HandleFunc("/complete-todo/{id:[0-9]+}", a.completeTodoHandler).Methods("GET")

	return *a

}

func (a *AppHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "./todo.html", http.StatusTemporaryRedirect)
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
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	isSuccess := a.db.RemoveTodo(id)

	rd.JSON(w, http.StatusOK, Success{IsSuccess: isSuccess})
}

func (a *AppHandler) completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	complete := r.FormValue("complete") == "true"

	isSuccess := a.db.CompleteTodo(id, complete)

	rd.JSON(w, http.StatusOK, Success{IsSuccess: isSuccess})

}

func (a *AppHandler) Close() {
	a.db.Close()
}
