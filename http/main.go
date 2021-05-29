package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// API 호출 테스트
func main() {
	// Request 객체 생성
	req, err := http.NewRequest("GET", "http://1.246.219.218:8082/api/product", nil)
	if err != nil {
		panic(err)
	}

	//필요시 헤더 추가 가능
	req.Header.Add("accept", "application/json")
	req.Header.Add("XAPIKEY", "eyJhbGciOiJIUzI1NiJ9.e30.ZsNiImZZERkZDt7W_5TKDmGc518hh0avtgZzTuOiOvQ==")

	// Client객체에서 Request 실행
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 결과 출력
	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes) //바이트를 문자열로
	fmt.Println(str)
}
