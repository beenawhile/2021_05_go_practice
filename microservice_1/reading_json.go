package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type helloWorldResponse struct {
	// 출력 필드를 "message"로 바꿈
	Message string `json:"message"`
	// 이 필드는 출력하지 않음
	Author string `json:"-"`
	// 값이 비어 있으면 필드를 출력하지 않음
	Date string `json:",omitempty"`
	// 출력을 문자열로 변환하고 이름을 "id"로 변경
	Id int `json:"id,string"`
}

// 채널, 복소수 및 함수는 json으로 인코딩할 수 없음
// 순환 데이터 구조를 나타낼 수 없음
// indent로 깔끔하게 양식을 지정한 JSON을 보내려면 MarshalIndent 사용하면 됨

func jsonMarshalHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello World", Id: 2, Author: "Kim"}
	data, err := json.Marshal(response)
	if err != nil {
		panic("Ooops!")
	}

	fmt.Fprint(w, string(data))
}

// 위의 방법은 구조체를 바이트 배열로 디코딩한 후 응답스트림에 쓰는 방법 => 느리다
// => 직접 스트림에 쓸 수 있는 encoder, decoder가 효율적

func jsonEncoderHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello World", Id: 2, Author: "Kim"}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}

// 두개 benchmark test
