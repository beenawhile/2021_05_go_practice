// thread
// - 현대 프로그래밍 언어에서 가장 중요한 개념 중 하나
// - 가장 많이 쓰이는 것 중의 하나
// - 가장 까다로운 문제가 생기는 것 중의 하나
// - 사전적 의미 : 실, 줄, 시냇물
// - 튜링머신을 상상했을 때 튜링머신이 읽는 문서 뭉치 1개 : thread
// - 이전에 cpu 코어가 하나 일때는 어떻게 multi-tasking 처리를 했는가?
//  => thread를 매우 빠르게 전환하면서 마치 multi tasking 하는 것처럼 보이게 함
// - 이러한 thread 전환을 context switching이라고 함. 전환에는 전환비용이 듬.
//  => context switching이 자주 발생하게 되면 전환하는데 비용이 더 커지기 때문에 효율이 떨어짐
//    - context switching이 안일어나게 하려면 어떻게 해야할까?
//		- context switching 이라는 것은 cpu 갯수 < thread 갯수 일 때 발생함
//		- 해결방법 : thread 갯수를 최대한 cpu 갯수에 가깝게 맞추면 된다!

// Go's Thread
//  - os가 기본 제공하는 thread : kernel thread
//  - go는 kernel thread를 잘게 잘라서 한번 포장(wrapping)해서 만들었음 => go thread
//  - 프로그래머는 go thread만 사용하면 kernel thread를 알 필요가 없게 됨
//  - 이렇게 만든 이유? => go는 context switching의 비용을 최대한 줄이기 위해 만들었기 때문에
//  - go thread 방식
//    - wrapping(OS thread 최소한 사용 + 필요한 thread) => NM thread
//    - NM thread : 하나의 OS thread(=N)에 여러개의 go thread(=M)가 들어갈 수 있음
//  - 프로그래머 입장에서 cpu 갯수, thread 갯수, context switching을 고려하지 않고 thread를 만들어도 go 내부에서 os thread 내부에 알아서 할당해주기 때문에 큰 걱정없이 코드 작성할 수 있음

// program, process, thread 차이
// - program >= process >= thread
// 1. program : 실행파일 + data
//  - 프로그램 구동 과정 : 프로그램 더블클릭 -> 실행파일을 메모리에 적재
// 2. process : 메모리에 올라가 있는 실행파일
//  - Thread 여러개를 가질 수 있음

// multi-thread 문제
//  - 데이터 동기화(synchronization)
package main

import (
	"fmt"
	"time"
)

// 출력이 섞여서 나오는 것을 알 수 있음
func threadTest() {
	go fun1()
	go fun1()
	for i := 0; i < 20; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("main :", i)
	}
	fmt.Scanln()
}

func fun1() {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("fun1:", i)
	}
}
