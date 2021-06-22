package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Helo world!"))
	})
	http.ListenAndServe(":9999", nil)
}
