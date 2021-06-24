package handlers

import (
	"encoding/json"
	"microservice_4/data"
	"net/http"
)

type SearchHandler struct {
}

func (s *SearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

type searchRequest struct {
	Query string `json:"query"`
}

type searchResponse struct {
	Kittens []data.Kitten `json:"kittens"`
}

type Search struct {
	DataStore data.Store
}

func (s *Search) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	request := new(searchRequest)
	err := decoder.Decode(request)
	if err != nil || len(request.Query) < 1 {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
		return
	}

	kittens := s.DataStore.Search(request.Query)

	encoder := json.NewEncoder(rw)
	encoder.Encode(searchResponse{Kittens: kittens})

}
