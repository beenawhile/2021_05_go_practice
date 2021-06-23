package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func corsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		// * : 모든 API 와 상호작용할 수 있음
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET")
		// OPTIONS 요청에 대해 본문을 리턴하는 것이 유효하지 않으므로 204 No Content를 보냄
		w.WriteHeader(http.StatusNoContent)
		return
	}

	response := helloWorldResponse{Message: "Hello world"}
	data, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(data))
}
