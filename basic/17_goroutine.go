package main

import (
	"fmt"
	"sync"
	"time"
)

// goroutine
// - Go runtime이 관리하는 lightweight 논리 쓰레드
// - OS 쓰레드보다 훨씬 가볍게 비동기 Concurrent 처리를 구현하기 위하여 만든 것
// - 기본적으로 Go 런타임이 자체 관리
// - go 키워드를 사용하여 함수를 호출하면, 런타임시 새로운 goroutine을 실행
// - 비동기적으로 함수루틴을 실행하므로, 여러 코드를 동시에 실행하는데 사용
func goroutineTest() {
	// 동기적
	saySomething("Sync")
	// 비동기적 멀티 스레드(OS 스레드 보다 크기가 작아 효율적)
	go saySomething("Async1")
	go saySomething("Async2")
	go saySomething("Async3")

	time.Sleep(time.Second * 3)

	// 익명함수 goroutine

	// WaitGroup 생성. 2개의 Go루틴을 기다림
	// WaitGroup - Go루틴들이 끝날 때까지 기다리는 역할
	var wait sync.WaitGroup
	wait.Add(2)

	// 익명함수를 사용한 goroutine
	go func() {
		defer wait.Done()
		fmt.Println("Hello")
	}()

	go func(msg string) {
		defer wait.Done()
		fmt.Println(msg)
	}("Hi")

	// Wait - Go루틴들이 모두 끝날때 까지 대기
	wait.Wait()

	fmt.Println("Done")

	// 다중 CPU 처리
	// - go는 default로 1개의 CPU 사용 => 1개의 작업을 시분할하여 처리(=Concurrent processing)
	// - 머신이 복수개 CPU를 가진 경우, Go 프로그램을 다중 CPU에서 병렬처리 하게 할 수 있음(=Parallel processing)
	// - runtime.GOMAXPROCS(CPU수)를 호출하여야 함

	// Concurrency vs Parallel
	// 	-	In programming, concurrency is the composition of independently executing processes,
	// 		while parallelism is the simultaneous execution of (possibly related) computations.
	// 		Concurrency is about dealing with lots of things at once.
	// 		Parallelism is about doing lots of things at once.

}

func saySomething(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}
