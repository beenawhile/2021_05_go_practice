package main

import (
	"encoding/json"
	"net/http"
)

type validationHanlder struct {
	next http.Handler
}

func (h validationHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	h.next.ServeHTTP(w, r)
}

func newValidationHandler(next http.Handler) http.Handler {
	return validationHanlder{next: next}
}

type helloWorldHandler struct{}

func newHelloWorldHandler() http.Handler {
	return helloWorldHandler{}
}

func (h helloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello"}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
