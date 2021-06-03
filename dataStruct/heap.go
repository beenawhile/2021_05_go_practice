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

package dataStruct

import "fmt"

type MaxHeap struct {
	list []int
}

func (h *MaxHeap) Push(v int) {
	// 맨마지막 자식 노드에 넣어줌
	h.list = append(h.list, v)

	idx := len(h.list) - 1
	// parent를 돌면서 새로 추가한 값이 더 큰지 확인
	for idx >= 0 {
		parentIdx := (idx - 1) / 2
		if parentIdx < 0 {
			break
		}
		if h.list[idx] > h.list[parentIdx] {
			// 새로 추가한 값이 크면 바꿔주기
			h.list[idx], h.list[parentIdx] = h.list[parentIdx], h.list[idx]
			idx = parentIdx
		} else {
			// 크지 않은 경우
			break
		}
	}
}

func (h *MaxHeap) Print() {
	fmt.Println(h.list)
}

func (h *MaxHeap) Pop() int {
	if len(h.list) == 0 {
		return 0
	}

	top := h.list[0]

	last := h.list[len(h.list)-1]

	// 맨 마지막 잘라냄
	h.list = h.list[:len(h.list)-1]

	if len(h.list) == 0 {
		return 0
	}
	// 맨 마지막 것을 처음으로 올림
	h.list[0] = last
	idx := 0
	for idx < len(h.list) {

		// 바꿨다는 것을 표시하기 위해 사용할 index
		swapIdx := -1

		leftIdx := idx*2 + 1

		if leftIdx >= len(h.list) {
			break
		}

		// left 비교해서 잠재적으로 바꿀 idx로 넣어줌
		if h.list[leftIdx] > h.list[idx] {
			swapIdx = leftIdx
		}

		// right 비교
		rightIdx := idx*2 + 2
		if rightIdx < len(h.list) {
			// 자식 노드가 더 있다는 의미
			if h.list[rightIdx] > h.list[idx] {
				if swapIdx < 0 || (swapIdx >= 0 && h.list[swapIdx] < h.list[rightIdx]) {
					swapIdx = rightIdx
				}
			}
		}

		if swapIdx <= 0 {
			break
		} else {
			h.list[idx], h.list[swapIdx] = h.list[swapIdx], h.list[idx]
			idx = swapIdx
		}
	}

	return top
}

type MinHeap struct {
	list []int
}

func (h *MinHeap) Push(v int) {
	// 맨마지막 자식 노드에 넣어줌
	h.list = append(h.list, v)

	idx := len(h.list) - 1
	// parent를 돌면서 새로 추가한 값이 더 큰지 확인
	for idx >= 0 {
		parentIdx := (idx - 1) / 2
		if parentIdx < 0 {
			break
		}
		if h.list[idx] < h.list[parentIdx] {
			h.list[idx], h.list[parentIdx] = h.list[parentIdx], h.list[idx]
			idx = parentIdx
		} else {
			break
		}
	}
}

func (h *MinHeap) Print() {
	fmt.Println(h.list)
}

func (h *MinHeap) Pop() int {
	if len(h.list) == 0 {
		return 0
	}

	top := h.list[0]
	last := h.list[len(h.list)-1]

	h.list = h.list[:len(h.list)-1]

	if len(h.list) == 0 {
		return 0
	}

	h.list[0] = last
	idx := 0
	for idx < len(h.list) {

		swapIdx := -1

		leftIdx := idx*2 + 1

		if leftIdx >= len(h.list) {
			break
		}

		if h.list[leftIdx] < h.list[idx] {
			swapIdx = leftIdx
		}

		rightIdx := idx*2 + 2
		if rightIdx < len(h.list) {
			if h.list[rightIdx] < h.list[idx] {
				if swapIdx < 0 || (swapIdx >= 0 && h.list[swapIdx] > h.list[rightIdx]) {
					swapIdx = rightIdx
				}
			}
		}

		if swapIdx <= 0 {
			break
		} else {
			h.list[idx], h.list[swapIdx] = h.list[swapIdx], h.list[idx]
			idx = swapIdx
		}
	}

	return top
}

func (h *MinHeap) Count() int {
	return len(h.list)
}
