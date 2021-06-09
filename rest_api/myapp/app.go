package myapp

import (
	"encoding/json"
	"fmt"

	"html/template" // 특수문자가 탈락되서 나옴
	"net/http"
	"os"
	"strconv"

	// "text/template" // 특수문자가 탈락되어 나오지 않음
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

// update용 struct
// type UpdateUser struct {
// 	ID               int `json:"id"`
// 	UpdatedFirstName bool
// 	FirstName        string `json:"first_name"`
// 	UpdatedLastName  bool
// 	LastName         string `json:"last_name"`
// 	UpdatedEmail     bool
// 	Email            string    `json:"email"`
// 	CreatedAt        time.Time `json:"created_at"`
// }

// user 정보를 담고 있는 map
var userMap map[int]*User
var lastID int

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello, world")
}

// 이렇게 만든 method를 template에서 사용할 수 있음 => tmpl1.tmpl 참고
func (u User) hasLongName() bool {
	return len(u.FirstName) > 5
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	user := User{ID: 3, FirstName: "Tuckerister", LastName: "Kim", Email: "tucker@naver.com"}
	user2 := User{ID: 4, FirstName: "aaa", LastName: "Lee", Email: "aaa@naver.com"}
	// 3.list로 넘겼을 때 => 파일에 range 부분 확인
	users := []User{user, user2}
	// 1.코드안에 template을 만들었을 때
	// tmpl, err := template.New("Tmpl1").Parse("Name : {{.FirstName}} {{.LastName}}\nEmail : {{.Email}}\n")
	tmpl, err := template.New("Tmpl1").ParseFiles("templates/tmpl1.tmpl", "templates/tmpl2.tmpl")
	if err != nil {
		panic(err)
	}
	// 1.코드안에 template을 만들었을 때
	// tmpl.Execute(os.Stdout, user)
	// tmpl.Execute(os.Stdout, user2)

	// 2.template 파일로 받았을 때
	// tmpl.ExecuteTemplate(os.Stdout, "tmpl1.tmpl", user)
	// tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", user2)

	// 3.list로 넘겼을 때 => 파일에 range 부분 확인
	tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", users)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	if len(userMap) == 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No Users")
		return
	}

	users := []*User{}
	for _, u := range userMap {
		users = append(users, u)
	}

	data, _ := json.Marshal(users)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
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

func deleteUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	_, ok := userMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User Id:", id)
		return
	}

	delete(userMap, id)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted User Id:", id)

}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	updateUser := new(User)
	err := json.NewDecoder(r.Body).Decode(updateUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	user, ok := userMap[updateUser.ID]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User Id:", updateUser.ID)
		return
	}

	if updateUser.FirstName != "" {
		user.FirstName = updateUser.FirstName
	}

	// 일부러 빈값으로 바꾸고 싶을 때는 다음 로직을 거치지 않는 문제가 생김
	// => 실무에서는 update 용 struct를 따로 만든다고 함
	if updateUser.LastName != "" {
		user.LastName = updateUser.LastName
	}

	if updateUser.Email != "" {
		user.Email = updateUser.Email
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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
	mux.HandleFunc("/users", updateUserHandler).Methods("PUT")
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler).Methods("GET")
	mux.HandleFunc("/users/{id:[0-9]+}", deleteUserInfoHandler).Methods("DELETE")

	// template
	mux.HandleFunc("/template", templateHandler)

	// 따로 지정하지 않으면 다른 uri 가 들어오더라도 테스트 통과하게 됨
	// 예외 처리

	return mux
}
