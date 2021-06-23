package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func fetchGoogle(t *testing.T) {
	r, _ := http.NewRequest("GET", "https://google.com", nil)

	timeoutRequest, cancelFunc := context.WithTimeout(r.Context(), 1*time.Millisecond)
	defer cancelFunc()

	r = r.WithContext(timeoutRequest)

	_, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// context 패키지를 통해 go routine이 객체에 안전하게 접근하도록 구현하는 방법
type validation2Hanlder struct {
	next http.Handler
}

type validationContextKey string

func newValidation2Handler(next http.Handler) http.Handler {
	return validation2Hanlder{next: next}
}

func (h validation2Hanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// validationContextKey을 써서 단순 문자열을 넘기지 않은 이유?
	//  - context가 패키지를 넘어 전달되는 경우가 많기 때문. 만약 단순 문자열을 사용한다면 키 충돌을 일으킬 수 있음
	//  - 우리가 조작할 수 없으면서 컨텍스트를 사용하는 또 다른 패키지 역시 name이라는 키에 값을 쓰려고 하는 경우 두 번째로 값을 쓰는 패키지가 의도치 않게 컨텍스트 값을 덮어쓰게 됨
	c := context.WithValue(r.Context(), validationContextKey("name"), request.Name)
	r = r.WithContext(c)

	h.next.ServeHTTP(w, r)

}

type helloWorld2Handler struct{}

func newHelloWorld2Handler() http.Handler {
	return helloWorld2Handler{}
}

func (h helloWorld2Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.Context().Value(validationContextKey("name")).(string)
	response := helloWorldResponse{Message: "Hello " + name}

	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}
