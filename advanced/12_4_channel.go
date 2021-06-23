// channel : go에서 제공하는 데이터 타입. 일종의 queue.
//  - multi thread 환경에서 dead lock을 해결할 수 있는 방법
//  - thread safe
//  - fixed size
//  - 선언
//     var 이름 chan 데이터타입
//  - 초기화
//     이름 := make(chan 데이터타입, (크기)), 크기 지정하지 않으면 0 => 하나도 집어넣을 수 있는 것이 아니라, 어딘가에서 빼주지 않으면 넣는 연산이 끝나지 않음
//  - channel은 slice처럼 자동으로 늘어나지 않고 크기는 고정됨
//  - queue 처럼 push, pop 지원. <- 화살표 연산자로 사용
//   push : a <- 10
//   pop : b:= <- a

package main

import "fmt"

func channelTest() {
	var c chan int
	c = make(chan int, 1)

	c <- 10
	v := <-c
	fmt.Println(v)

	// channel size = 0 인 경우
	c1 := make(chan int)
	go pop(c1) // 빼주기 위해 추가
	c1 <- 20   // 10 을 넣을 때 channel size = 0 이라서 멈춤. 다른 쪽에서 빼줘야함

	fmt.Println("End of program")

	// * 이런 channel을 이용해서 conveyor belt pattern (=producer-consumer pattern) 을 구축할 수 있음
	// - 13_4 에서 car factory 예시를 만들어 봄
}

func pop(c chan int) {
	fmt.Println("Pop function")
	v := <-c // 여기서 대기 하고 있다가 다른 thread에서 값을 보내면 그것을 받아서 진행
	fmt.Println(v)
}
