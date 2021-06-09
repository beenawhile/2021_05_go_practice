package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/antage/eventsource"
	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
)

func main() {
	mux := pat.New()

	// event source 라이브러리 추가하기
	// go get github.com/antage/eventsource
	// 메세지 채널 초기화
	msgCh = make(chan Message)

	es := eventsource.New(nil, nil)
	defer es.Close()

	// message 채널 열어주기
	go processMsgCh(es)

	mux.Post("/messages", postMessageHandler)
	mux.Handle("/stream", es)

	mux.Post("/users", addUserHandler)
	mux.Delete("/users", leftUserHandler)

	n := negroni.Classic()

	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}

func postMessageHandler(w http.ResponseWriter, r *http.Request) {
	msg := r.FormValue("msg")
	name := r.FormValue("name")
	sendMessage(name, msg)
}

// multi thread 환경에서 돌아가고 있으니
// thread 추가하여 queue 형태로 보내는 것이 더 좋음

type Message struct {
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

var msgCh chan Message

func sendMessage(name, msg string) {
	// send message to all clients
	msgCh <- Message{Name: name, Msg: msg}
}

func processMsgCh(es eventsource.EventSource) {
	for msg := range msgCh {
		data, _ := json.Marshal(msg)
		es.SendEventMessage(string(data), "", strconv.Itoa(time.Now().Nanosecond()))
	}
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("name")
	sendMessage("", fmt.Sprintf("added user: %s", username))
}

func leftUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	sendMessage("", fmt.Sprintf("left user: %s", username))
}
