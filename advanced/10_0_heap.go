// heap
//  - 최대값 또는 최소값을 빠르게 찾기위해서 고안된 트리모양의 자료구조
// 1. max heap : 부모 노드가 자식 노드보다 항상 크거나 같아야하는 트리모양 자료 구조
// 2. min heap : 부모 노드가 자식 노드보다 항상 작거나 같아야하는 트리모양 자료 구조

// 장점
//  1. 최댓값, 최솟값 찾을 때 좋음
//  2. Priority queue(우선순위 큐)를 heap으로 만들 수 있음
//    ※ Priority queue : 우선순위가 높은 것이 먼저 나오는 queue (ex. 응급실 대기열을 만든다고 생각했을 때 응급한 사람이 먼저 진료받는 개념)
//  3. 정렬을 만들 수 있음

// push & pop 속도 비교
//  - push : O(log2N)
//  - pop : O(log2N)

// heap 정렬 속도
//  - push + pop : O(2*N*log2N) + O(2*N*log2N) = O(2*N*log2N)
//  - big O 법에서 크기가 매우 커지면 상수는 의미 없기 때문에 O(N*log2N)

// 구현하기 위해서는 맨 뒤 노드가 무엇인지 알고 있어야함
// - linked list로는 어려움
// - array로 많이 구현함
// - array 안에 맨 위 값 부터 순서대로 넣어서 꺼내는 방식
// - list 안에서 각 자식 노드는 어떻게 찾는가?
// - n 번째 index의 left : 2n+1
// - n 번째 index의 right : 2n+2
// - n 번째 index의 parent : (n-1)/2

package main

import (
	"dataStruct"
	"fmt"
)

func heapTest() {
	h := dataStruct.MaxHeap{}
	h.Push(2)
	h.Push(7)
	h.Push(9)
	h.Push(8)
	h.Push(5)
	h.Push(6)
	h.Push(6)

	h.Print()

	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())

}
