package main

// 1. 그간 사용했던 gorila의 mux 말고 pat을 써보자
// $go get github.com/gorilla/pat

// 2. api 호출 테스트 하는 것이 귀찮다? => 패키지 가져오기
// - render(unrolled)
// $go get github.com/unrolled/render

// 3. 기본적으로 많이 쓰이는 부가기능을 제공하는 패키지
// - negroni
// - recovery, logger, static file serving
// file serving handler 만들었던거 기억해보기
// $go get github.com/urfave/negroni

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/pat"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

// render 전역변수 추가
var rd *render.Render

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	// template 확장자를 tmpl 외의 다른 것을 사용하고 싶을 때
	// template 폴더도 다른 것 사용하고 싶을 때 설정
	rd = render.New(render.Options{
		Directory:  "template",
		Extensions: []string{".html", ".tmpl"},
		// html 여러개로 layout 짜는 경우
		Layout: "hello"},
	)
	// rd = render.New()

	// mux := mux.NewRouter()
	mux := pat.New()

	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserHandler)

	mux.Get("/hello", helloHandler)

	// 기본 파일 서버
	// mux.Handle("/", http.FileServer(http.Dir("public")))

	// negroni 확장 기능
	// log 기능, static file server 기능
	n := negroni.Classic()
	n.UseHandler(mux)

	// http.ListenAndServe(":3000", mux)
	http.ListenAndServe(":3000", n)

}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {

	user := User{Name: "tucker", Email: "tucker@naver.com"}

	// rd로 다음과 같이 바꿈
	rd.JSON(w, http.StatusOK, user)

	// w.Header().Add("Content-type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// data, _ := json.Marshal(user)
	// fmt.Fprint(w, string(data))

}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		// fmt.Fprint(w, err)
		rd.Text(w, http.StatusBadRequest, err.Error())
		return
	}
	user.CreatedAt = time.Now()
	// w.Header().Add("Content-type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// data, _ := json.Marshal(user)
	// fmt.Fprint(w, string(data))

	// rd로 다음과 같이 바꿈
	rd.JSON(w, http.StatusOK, user)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "tucker", Email: "tucker@naver.com"}
	// tmpl, err := template.New("Hello").ParseFiles("templates/hello.tmpl")
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	fmt.Fprint(w, err)
	// 	return
	// }

	// tmpl.ExecuteTemplate(w, "hello.tmpl", "Tucker")

	// 참고하는 template은 1. templates/ 아래, 2. .tmpl 이라는 확장자를 가져야함
	// rd.HTML(w, http.StatusOK, "hello", "Tucker")
	// 확장자가 바뀐 경우? => render 초기에 설정할 때 option 값을 넣어줘야함
	rd.HTML(w, http.StatusOK, "body", user)

}
