package app

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"todos/model"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

// 암호화되는 cookie store를 만듬
var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
var rd *render.Render = render.New()

// 모든 엔드포인트에 대해서 decorator pattern으로 login 체크 확인하는 것 추가

// 엔드포인트 접속 될 때 세션 아이디를 읽어옴
var getSessionID = func(r *http.Request) string {
	// session cookie에 id 저장
	session, err := store.Get(r, "session")
	if err != nil {
		return ""
	}

	val := session.Values["id"]
	if val == nil {
		return ""
	}
	return val.(string)
}

type AppHandler struct {
	// interface를 포함하고 있다 => 상속하고 다른 의미 (has 관계이지 is 관계가 아님)
	http.Handler
	db model.DBHandler
}

func MakeHandler(filepath string) *AppHandler {

	// ---- Decorator pattern ----

	r := mux.NewRouter()

	// login check middle ware 추가
	// n := negroni.Classic()
	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger(), negroni.HandlerFunc(CheckSignin), negroni.NewStatic(http.Dir("public"))) // 체인 형태로 decorator가 물려있음
	//
	n.UseHandler(r)

	a := &AppHandler{
		Handler: n,
		db:      model.NewDBHandler(filepath),
	}
	// ---------------------------

	r.HandleFunc("/", a.indexHandler)
	r.HandleFunc("/todos", a.getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos", a.addTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", a.removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/complete-todo/{id:[0-9]+}", a.completeTodoHandler).Methods("GET")
	r.HandleFunc("/auth/google/login", googleLoginHandler)
	r.HandleFunc("/auth/google/callback", googleAuthCallback)

	return a
}

func CheckSignin(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// 주의할점 : signin 요청인지 체크 안하면 signin 페이지에서 무한 루프를 돌기 때문에 체크 해줘야함
	if strings.Contains(r.URL.Path, "/signin") || strings.Contains(r.URL.Path, "/auth") {
		next(w, r)
		return
	}

	// 1. signedin  => next
	// 2. signedin  => redirect to signin page
	sessionID := getSessionID(r)
	if sessionID != "" {
		next(w, r)
		return
	}
	http.Redirect(w, r, "/signin.html", http.StatusTemporaryRedirect)
}

func (a *AppHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func (a *AppHandler) getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	sessionId := getSessionID(r)
	list := a.db.GetTodos(sessionId)
	rd.JSON(w, http.StatusOK, list)
}

func (a *AppHandler) addTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	sessionId := getSessionID(r)

	todo := a.db.AddTodo(name, sessionId)

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
