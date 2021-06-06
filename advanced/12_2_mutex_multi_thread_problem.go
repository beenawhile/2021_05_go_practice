package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// multi-thread 문제
//  - 데이터 동기화(synchronization)

// - 지금 multi threading 은 어떤 방식인지 이해하기
//  - 종이에 그림을 그릴 때 한 종이에다 범위는 신경쓰지 않고 동시에 그리는 방법

// 어느 정도 범위로 locking을 해야할지는 매우 복잡한 문제가 됨
// lock이 무조건 좋지 않고 dead lock이 발생할 수 있음
// channel을 사용하여 일정 부분 문제를 해결할 수 있음. 하지만 channel이 정답은 아님!
// channel과 lock이 다른 접근 방법이기 때문에 각각이 필요할 때 각각의 상황에 맞춰 써야함
// lock 을 사용할 때는 주의해야할 점이 많다는 것만 기억하자

// dead lock : dead end(막다른 길) 처럼 lock이 막혔다는 의미
// - multi-thread programming을 할 때 발생할 수 있는 치명적 오류

// dead lock 문제를 해결할 수 있는 방법
// => 공장 컨베이어 벨트 방식 : 각자 자기 차례가 왔을 때 할 수 있는 부분만 하고 넘기는 방식. 속도면에서 효율적.
//  = 생산자 - 소비자 패턴(producer consumer pattern)
// - go 에서는 이것을 쉽게 하기 위해 channel keyword을 제공
//  => channel은 일종의 queue(FIFO)

type Account struct {
	balance int
	// 문제 해결 방법 : mutex
	mutex *sync.Mutex
}

func (a *Account) Withdraw(val int) {
	a.mutex.Lock()
	a.balance -= val
	a.mutex.Unlock()
}

func (a *Account) Deposit(val int) {
	a.mutex.Lock()
	a.balance += val
	a.mutex.Unlock()
}

func (a *Account) Balance() int {
	a.mutex.Lock()
	balance := a.balance
	a.mutex.Unlock()
	return balance
}

var accounts []*Account
var globalLock *sync.Mutex

func Transfer(sender, receiver int, money int) {
	globalLock.Lock()
	accounts[sender].Withdraw(money)
	accounts[receiver].Deposit(money)
	globalLock.Unlock()

}

func GetTotalBalance() int {
	globalLock.Lock()
	total := 0
	for i := 0; i < len(accounts); i++ {
		total += accounts[i].Balance()
	}
	globalLock.Unlock()
	return total
}

func RandomTransfer() {
	var sender, balance int
	for {
		sender = rand.Intn(len(accounts))
		balance = accounts[sender].Balance()
		if balance > 0 {
			break
		}
	}

	var receiver int
	for {
		receiver = rand.Intn(len(accounts))
		if sender != receiver {
			break
		}
	}

	money := rand.Intn(balance)
	Transfer(sender, receiver, money)

}

func GoTransfer() {
	for {
		RandomTransfer()
	}
}

func PrintTotalBalance() {
	fmt.Printf("Total : %d\n", GetTotalBalance())
}

func mutexMultiThreadProblem() {
	for i := 0; i < 20; i++ {
		accounts = append(accounts, &Account{balance: 1000, mutex: &sync.Mutex{}})
	}

	globalLock = &sync.Mutex{}

	PrintTotalBalance()

	// go routine이 여러개 생기면서 memory를 흐트려 놓기 때문에 balance가 바뀌게 됨
	for i := 0; i < 10; i++ {
		go GoTransfer()
	}

	for {
		PrintTotalBalance()
		time.Sleep(100 * time.Millisecond)
	}

	// 결과
	// Total : 20000
	// Total : 20000
	// Total : 8213
	// Total : 18381
	// Total : 11096
	// Total : 19890
	// Total : 21974
	// Total : 32647
	// Total : 22405
	// Total : 31045
	// Total : 70160
	// Total : 175385
	// Total : 78025
	// 결과가 바뀌는 이유? multi-thread
	// a.balance -= val 은 하나의 실행문이 아님
	// => a.balance = a.balance - val
	// 이 문장 역시 assembly language로 바꾸면 여러개의 문장이 될 수 있음
	// 다음과 비슷한 코드가 될것임 (pseudo code)
	// --------------------------
	// Load RegA a.balance
	// Load RegB val
	// Del RegA RegB , &a.balance
	// --------------------------
	// 다른 cpu들에서 같은 연산을 하다보면 값이 꼬이게 됨

	// 어떻게 문제를 해결하는가?
	// 1. Lock : 가장 대표적인 방법 == Mutex
	//  - thread 어떤 자원에 접근했을 때 잠시 잠궜다가 사용이 끝났을 때 풀어줌

}
