package main

import (
	"fmt"
	"strconv"
	"time"
)

type Car struct {
	val string
}

func MakeTire(carChan chan Car, outChan chan Car) {
	for {

		car := <-carChan
		car.val += "Tire, "
		outChan <- car
	}
}

func MakeEngine(carChan chan Car, outChan chan Car) {
	for {

		car := <-carChan
		car.val += "Engine, "
		outChan <- car
	}
}

func StartWork(chan1 chan Car) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Car{val: "Car" + strconv.Itoa(i)}
		i++
	}
}

// lock을 잡을 필요 없다는것 잘 보기
func carFactory() {
	chan1 := make(chan Car)
	chan2 := make(chan Car)
	chan3 := make(chan Car)

	go StartWork(chan1)
	go MakeTire(chan1, chan2)
	go MakeEngine(chan2, chan3)

	for {
		result := <-chan3
		fmt.Println(result)
	}

}

// go thread, channel, select : 이 3개를 사용하면 효율적으로 multi thread를 만들 수 있음!

// multi thread는 복잡한데 왜 사용해야하는가? => machine의 성능을 최대화 하기 위해서
// - 공짜 점심은 끝났다로 표현하기도 함 => moor's law가 한계에 부딪히게 됬음 => multi core 만들어내기 시작했음
//  cpu가 더이상 2배로 빨라지지 않기 때문에 효율성이 중요함

// 실행단위 : cpu core에서 실행하는 하나의 단위. 프로세스와 스레드를 포괄하는 개념
// (부연설명 없는)프로세스 : 하나의 스레드만 가지고 있는 단일 스레드 프로세스
// 동시성 : 한 순간에 여러가지 일이 아니라, 짧은 전환으로 여러가지 일을 동시에 처리하는 것 처럼 보이는 것

// 프로그램이 프로세스가 되면서 일어나는 일
// 1. 프로세스가 필요로 하는 재료들이 메모리에 올라감
//  - code : 실행 명령을 포함하는 코드들
//  - data : static 변수 혹은 global 변수
//  - heap : 동적 메모리 영역
//  - stack : 지역변수, 매개변수, 반환 값 등 일시적 데이터
// 2. 해당 프로세스에 대한 정보를 담고 있는 PCB 블럭이 만들어짐
//  - PID, Program counter, 등등

// 하나의 프로세스가 CPU를 점유하고 있으면 다른 프로세스를 실행할 수 없음
// -> 동시에 여러개의 프로세스를 수행하기 위해 시분할(=짧은 텀 반복)로 전환해서 실행 => 동시성

// 동시성 과정
//  - PCB1실행, PCB2준비 -> PCB1준비, PCB2준비 -> PCB1준비, PCB2 실행 (이러한 과정을 context switching 이라 부름)
// 피곤하고 힘든 작업
// => 경량화된 프로세스 thread의 등장!

// thread : 공통으로 사용할 것들은 두고 일부 자원만 메모리에 올리고 내림 => 모두 다빼고 다 넣을 필요 없음 => 캐싱 적중률이 올라감

// multi process & multi thread
// - 주의할점 : 두 가지 모두 한 애플리케이션에 대한 처리방식의 일종
// - 여러 프로그램을 띄워놓은 것을 multi process라고 이해하면 안됨
// 1. multi process
//  - 한 애플리케이션이 여러가지 일을 처리할 때(여러 명의 로그인 처리)
//  - 한 프로세스는 하나의 일만 처리할 수 있고 동시처리 할 수 없음
//  - 해결하기 위해 부모 프로세스가 fork하여 자식 프로세스를 여러 개 생성하여 일 처리
//  - 이 때, 자식 프로세스는 부모와 별개의 메모리 영역을 확보하게 됨
// 2. thread
//  - 한 프로세스 내에서 구분이 지어진 실행 단위
//  - multi thread : 프로세스 내에서 분리해서 여러 스레드로 나뉘어서 실행단위가 나뉘어 진 것
//  - 한 애플리케이션에 대해 작업단위가 나누어 진 것 (ex. vs code intelisense, debugging, code duplication)

// 차이점
// 1. multi process
//  - 각 프로세스는 독립적
//  - IPC를 사용한 통신 해야함
//  - 자원 소모적, 개별 메모리 차지
//  - context switching 비용이 큼
//  - 동기화 작업이 필요하지 않음
// 2. multi thread
//  - thread 끼리 긴밀하게 연결되어 있음
//  - 공유된 자원으로 통신 비용 절감
//  - 공유된 자원으로 메모리가 효율적임
//  - context switching 비용이 적음
//  - 공유 자원 관리 해야함

// 그럼 효율적인 multi thread 말고 multi process는 왜 필요한가?
// - 예시로 보자
// - ie 를 사용할 때 잘못되면 ie 전체가 꺼지는 경우 생김 => multi thread를 사용하였기 때문
// - chrome에서는 탭간 독립적 => multi process를 사용하였기 때문
// ※ 다 장단점이 있음

// multi core
//  - 동시성, 병렬처리
//  - concurrency(동시성) : 하나의 코어에서 하나이상 프로세스(or 스레드)가 번갈아가면서 진행되지만 동시에 진행되는 것처럼 보이는 것
//  - parallelism(병렬처리) : 둘 이상의 코어에서 동시에 하나 이상의 프로세스(or 스레드)가 한꺼번에 진행되는 것
