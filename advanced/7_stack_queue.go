package main

import (
	"dataStruct"
	"fmt"
)

// Stack & Queue : 성격이 비슷하고 구현방법이 비슷함. 동작은 반대로

// Stack : FILO
// Queue : FIFO

func stackQueueTest() {

	// 1. slice로 만들기

	// - Stack
	stack := []int{}
	for i := 1; i < 6; i++ {
		stack = append(stack, i)
	}
	// Add : O(N) (cap 넘어갈 때)

	fmt.Println(stack)

	for len(stack) > 0 {
		var last int
		last, stack = stack[len(stack)-1], stack[:len(stack)-1]
		fmt.Println(last)
	}
	// Remove : O(1)

	// - Queue
	queue := []int{}
	for i := 1; i < 6; i++ {
		queue = append(queue, i)
	}
	// Add : O(N) (cap 넘어갈 때)

	fmt.Println(queue)

	var first int
	for len(queue) > 0 {

		first, queue = queue[0], queue[1:]
		fmt.Println(first)
	}
	// Remove : O(1)

	// 2. linked list로 만들기
	// - stack
	stack2 := dataStruct.NewStack()

	for i := 1; i < 6; i++ {
		stack2.Push(i)
	}

	fmt.Println("New Stack")

	for !stack2.Empty() {
		val := stack2.Pop()
		fmt.Printf("%d -> ", val)
	}
	// Add : O(1)
	// Remove : O(1)
	fmt.Println()

	// - Queue
	queue2 := dataStruct.NewQueue()
	for i := 1; i < 6; i++ {
		queue2.Push(i)
	}

	fmt.Println("New Queue")

	for !queue2.Empty() {
		val := queue2.Pop()
		fmt.Printf("%d -> ", val)
	}

}

// stack & queue 를 어디에 사용하는가?
// - C++에는 stack memory, heap memory 있음
// - for문 2번 돌때 늦게 정의된 안쪽 for문의 변수 부터 메모리에서 내려간다 => stack
// - queue : 대기열 순서대로 진행하게 하고 싶다 할 때 사용
