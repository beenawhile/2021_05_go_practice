package main

import (
	"fmt"
	"os"
)

// error
// - Go는 내장타입으로 error 라는 interface 타입을 갖음
// - Go Error는 이 error 인터페이스를 통해서 주고 받게 됨
// - 이 interface는 Error() string 이라는 하나의 메소드를 갖음
// - 개발자가 인터페이스를 구현하는 커스텀 에러 타입을 만들 수 있음

func errorTest() {
	// f, err := os.Open("C:\\temp\\1.txt")
	// if err != nil {
	// 	fmt.Println("Error Occurred !!!")
	// 	log.Fatal(err.Error())
	// }
	// fmt.Println(f.Name())

	// 다른 에러 처리방법
	_, err := os.Open("C:\\temp\\1.txt")
	switch err.(type) {
	default:
		fmt.Println("ok")
	case error:
		fmt.Println(err.Error())
	}

}

type error interface {
	Error() string
}
