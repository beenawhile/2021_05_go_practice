package main

import (
	"fmt"
	"log"
	"net/http"
	"todos/app"
)

const (
	HOST        = "192.168.0.11"
	PORT        = 5432
	DB_USER     = "yjpmes"
	DB_PASSWORD = "yjpmes"
	DB_NAME     = "yjpmes"
)

func main() {

	pgsqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, DB_USER, DB_PASSWORD, DB_NAME)

	mux := app.MakeHandler(pgsqlConn)
	// mux := app.MakeHandler("./test.db") // 파일 경로를 안에 넣지 않아야 할 때는 flag 패키지를 사용 (method : flag.Args())
	defer mux.Close()

	log.Println("Started App")
	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		panic(err)
	}
}
