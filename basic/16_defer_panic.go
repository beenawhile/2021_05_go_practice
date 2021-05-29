package main

import (
	"fmt"
	"os"
)

// defer
// - 특정 문장 혹은 함수를 나중에(defer를 호출하는 함수가 리턴하기 직전에) 실행하게 함
// - C#, Java 등에서 finally 블럭처럼 마지막에 clean-up 작업을 위해 사용됨
// - 파일을 open 한 후 close 하는 작업에 defer을 사용하기도 함

func deferTest() {
	f, err := os.Open("1.txt")
	if err != nil {
		panic(err)
	}
	// main 마지막에 파일 close 실행
	defer f.Close()

	// 파일 읽기
	bytes := make([]byte, 1024)
	f.Read(bytes)
	fmt.Println(len(bytes))
}

// panic
// - 현재 함수를 즉시 멈추고 현재 함수에 defer 함수들을 실행한 후 즉시 리턴
// - 마지막에는 프로그램이 에러를 내고 종료하게 됨

func panicTest() {
	// 잘못된 파일명을 넣음
	openFile("Invalid.txt")

	// openFile() 안에서 panic이 실행되면
	// 아래 println 문장은 실행 안됨
	println("Done")
}

// func openFile(fn string) {
// 	f, err := os.Open(fn)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()
// }

// recover
// - panic 함수에 의한 패닉상태를 다시 정상상태로 돌리는 함수

func recoverTest() {
	// 잘못된 파일명을 넣음
	openFile("Invalid.txt")

	// recover에 의해
	// 다음문장 실행됨
	fmt.Println("Done")
}

func openFile(fn string) {
	// defer 함수. panic 호출시 실행됨
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}
	}()

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}
