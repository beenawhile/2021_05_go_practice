package app

import (
	"net/http"
	"strconv"
	"todos_practice/model"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

type Success struct {
	Success bool `json:"success"`
}

var rd *render.Render = render.New()

// 포함관계
type AppHandler struct {
	http.Handler
	dbHandler model.DBHandler
}

func MakeHandler() *AppHandler {

	r := mux.NewRouter()

	a := &AppHandler{
		Handler:   r,
		dbHandler: model.NewDBHandler(),
	}

	r.HandleFunc("/", a.indexHandler)
	r.HandleFunc("/todos", a.getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos", a.addTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", a.deleteTodoHandler).Methods("DELETE")
	r.HandleFunc("/complete-todo/{id:[0-9]+}", a.completeTodoHandler).Methods("GET")

	return a
}

func (a *AppHandler) Close() {
	a.dbHandler.Close()
}

func (a *AppHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "./todo.html", http.StatusPermanentRedirect)
}

func (a *AppHandler) getTodoListHandler(w http.ResponseWriter, r *http.Request) {

	list := a.dbHandler.GetTodos()
	rd.JSON(w, http.StatusOK, list)
}
func (a *AppHandler) addTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")

	todo := a.dbHandler.AddTodo(name)
	rd.JSON(w, http.StatusCreated, todo)
}

func (a *AppHandler) deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	ok := a.dbHandler.RemoveTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}

}

func (a *AppHandler) completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("complete") == "true"
	ok := a.dbHandler.CompleteTodo(id, complete)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}

}
