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

}
