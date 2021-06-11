package app

// import (
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/gorilla/mux"
// 	"github.com/unrolled/render"
// )

// type Todo struct {
// 	ID        int       `json:"id"`
// 	Name      string    `json:"name"`
// 	Completed bool      `json:"completed"`
// 	CreatedAt time.Time `json:"created_at"`
// }

// type Success struct {
// 	Success bool `json:"success"`
// }

// var rd *render.Render

// var todoMap map[int]*Todo

// func MakeHandler() http.Handler {

// 	todoMap = make(map[int]*Todo)

// 	rd = render.New()

// 	m := mux.NewRouter()

// 	m.HandleFunc("/", indexHandler)
// 	m.HandleFunc("/todos", getTodoListHandler).Methods("GET")
// 	m.HandleFunc("/todos", addTodoHandler).Methods("POST")
// 	m.HandleFunc("/todos/{id:[0-9]+}", deleteTodoHandler).Methods("DELETE")
// 	m.HandleFunc("/complete-todo/{id:[0-9]+}", completeTodoHandler).Methods("GET")

// 	return m
// }

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	http.Redirect(w, r, "./todo.html", http.StatusPermanentRedirect)
// }

// func getTodoListHandler(w http.ResponseWriter, r *http.Request) {
// 	list := []*Todo{}
// 	for _, v := range todoMap {
// 		list = append(list, v)
// 	}
// 	rd.JSON(w, http.StatusOK, list)
// }
// func addTodoHandler(w http.ResponseWriter, r *http.Request) {
// 	name := r.FormValue("name")
// 	id := len(todoMap)
// 	todo := &Todo{ID: id, Name: name, Completed: false, CreatedAt: time.Now()}
// 	todoMap[id] = todo
// 	rd.JSON(w, http.StatusCreated, todo)
// }

// func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, _ := strconv.Atoi(vars["id"])
// 	if _, ok := todoMap[id]; ok {
// 		delete(todoMap, id)
// 		rd.JSON(w, http.StatusOK, Success{true})
// 	} else {
// 		rd.JSON(w, http.StatusOK, Success{false})

// 	}
// }

// func completeTodoHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, _ := strconv.Atoi(vars["id"])
// 	complete := r.FormValue("complete") == "true"
// 	if todo, ok := todoMap[id]; ok {
// 		todo.Completed = complete
// 		rd.JSON(w, http.StatusOK, Success{true})
// 	} else {
// 		rd.JSON(w, http.StatusOK, Success{false})
// 	}
// }
