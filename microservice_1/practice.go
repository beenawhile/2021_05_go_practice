package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type validationKey string

type oneHandler struct {
	next http.Handler
}

func newOneHandler(next http.Handler) http.Handler {
	return oneHandler{next: next}
}

func (h oneHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), validationKey("name"), request.Name)
	r = r.WithContext(c)

	h.next.ServeHTTP(w, r)
}

type twoHandler struct{}

func newTwoHandler() http.Handler {
	return twoHandler{}
}

func (h twoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.Context().Value(validationKey("name")).(string)
	response := helloWorldResponse{Message: "Hello " + name}

	encoder := json.NewEncoder(w)
	encoder.Encode(&response)

}
