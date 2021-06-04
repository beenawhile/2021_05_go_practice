package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Account2 struct {
	balance int
	mutex   *sync.Mutex
}

func (a *Account2) Withdraw(val int) {
	// dead lock 피하기 위한 방법 1 : 작게 lock
	// a.mutex.Lock()
	a.balance -= val
	// a.mutex.Unlock()

}

func (a *Account2) Deposit(val int) {
	// dead lock 피하기 위한 방법 1 : 작게 lock
	// a.mutex.Lock()
	a.balance += val
	// a.mutex.Unlock()

}

func (a *Account2) Balance() int {
	balance := a.balance
	return balance
}

var accounts2 []*Account2

// dead lock 피하기 위한 방법 2 : 크게 lock
var globalLock2 *sync.Mutex

func Transfer2(sender, receiver int, money int) {

	// accounts2[sender].mutex.Lock()
	// accounts2[receiver].mutex.Lock()

	// 여기서 부터 진행이 안됨 => Dead Lock
	//  - lock 이 막혔다는 의미
	//  - 철학자의 식사시간 문제가 대표적으로 dead lock을 설명하는 방법

	// dead lock 은 매우 복잡한 문제
	// dead lock 을 잡으려면
	// 1. 작게 lock을 잡거나
	// 2. 크게 lock을 잡거나

	// dead lock 피하기 위한 방법 2 : 크게 lock

	globalLock2.Lock()

	accounts2[sender].Withdraw(money)
	accounts2[receiver].Deposit(money)

	globalLock2.Unlock()

	// accounts2[sender].mutex.Unlock()
	// accounts2[receiver].mutex.Unlock()

	fmt.Println("Transfer", sender, receiver, money)

}

func GetTotalBalance2() int {
	total := 0
	for i := 0; i < len(accounts2); i++ {
		total += accounts2[i].Balance()
	}
	return total
}

func RandomTransfer2() {
	var sender, balance int
	for {
		sender = rand.Intn(len(accounts2))
		balance = accounts2[sender].Balance()
		if balance > 0 {
			break
		}
	}

	var receiver int
	for {
		receiver = rand.Intn(len(accounts2))
		if sender != receiver {
			break
		}
	}

	money := rand.Intn(balance)
	Transfer2(sender, receiver, money)

}

func GoTransfer2() {
	for {
		RandomTransfer2()
	}
}

func PrintTotalBalance2() {
	fmt.Printf("Total : %d\n", GetTotalBalance2())
}

func deadLockMultiThreadProblem() {
	for i := 0; i < 20; i++ {
		accounts2 = append(accounts2, &Account2{balance: 1000, mutex: &sync.Mutex{}})
	}

	globalLock2 = &sync.Mutex{}

	go func() {
		for {
			Transfer2(0, 1, 100)
		}
	}()

	go func() {
		for {
			Transfer2(1, 0, 100)
		}
	}()

	for {
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
