package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Unmarshal : Marshal과 반대로 작동
// 필요에 따라 map, slice, pointer를 할당
// 대소문자를 구분하지는 않지만 정확하게 일치하는 것이 좋음

type helloWorldRequest struct {
	Name string `json:"name"`
}

func jsonUnmarshalHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var request helloWorldRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: "Hello " + request.Name}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

func jsonDecodeHandler(w http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: "Hello " + request.Name}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)

}
