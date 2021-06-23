package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	port := 8080

	http.HandleFunc("/json-marshal", jsonMarshalHandler)
	http.HandleFunc("/json-encoder", jsonEncoderHandler)
	http.HandleFunc("/json-unmarshal", jsonUnmarshalHandler)
	http.HandleFunc("/json-decoder", jsonDecodeHandler)

	cathandler := http.FileServer(http.Dir("./images"))
	// StripPrefix 를 수행하지 않으면 images/cat 디렉토리에서 이미지를 찾고 있을 것임
	// /cat/ 을 경로로 등록하므로 /cat, /cat/의 하위 디렉토리에 대한 모든 요청이 처리됨
	http.Handle("/cat/", http.StripPrefix("/cat/", cathandler))

	handler := newValidationHandler(newHelloWorldHandler())
	handler2 := newValidation2Handler(newHelloWorld2Handler())

	http.Handle("/hello", handler)
	http.Handle("/hello2", handler2)

	practiceHandler := newOneHandler(newTwoHandler())

	http.Handle("/practice", practiceHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
