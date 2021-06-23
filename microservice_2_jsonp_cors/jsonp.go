package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"`
}

func main() {
	port := 8080

	http.HandleFunc("/hellojsonp",
		http.HandlerFunc(jsonpHandler),
	)

	http.HandleFunc("/hellocors",
		http.HandlerFunc(corsHandler),
	)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}

func jsonpHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "HelloWorld"}
	data, err := json.Marshal(response)
	if err != nil {
		panic("Ooops")
	}

	callback := r.URL.Query().Get("callback")
	if callback != "" {
		r.Header.Set("Content-Type", "application/javascript")
		fmt.Fprintf(w, "%s(%s)", callback, string(data))
	} else {
		fmt.Fprint(w, string(data))
	}

}
