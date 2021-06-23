// select를 사용하여 무언가 작업을 하고 있다가 call 이 있으면 처리하고 다시 작업을 하는 예시

package main

import (
	"fmt"
	"time"
)

func push(c chan int) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		c <- i
		i++
	}
}

func selectCallTest() {

	c := make(chan int)

	go push(c)

	// 추가
	timerChan := time.After(10 * time.Second)
	tickTimer := time.Tick(2 * time.Second)

	for {
		select {
		// case v := <-c:
		// 	fmt.Println(v)
		// default:
		// 	fmt.Println("Idle")
		// 	time.Sleep(1 * time.Second)

		// time package 에는 시간 간격으로 발생하는 이벤트를 쉽게 구현하기 위해 Tick, After을 제공 => channel 을 제공 하는 것
		case v := <-c:
			fmt.Println(v)
		case <-timerChan:
			fmt.Println("timeout")
			return
			// 종료되지 않는다 => After method의 출력값으로 받는 것이 없기 때문에 계속 timer가 생성 초기화 되어서
			// timer를 따로 만들어 준다음	에 이를 참고하는 방법을 사용해야함
		case <-tickTimer:
			fmt.Println("Tick!")
		}
	}

}
