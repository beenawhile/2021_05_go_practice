// 어떤 m개 리스트가 주어질 때, 이 리스트에서 n 번째 큰 혹은 작은 값을 찾는 방법
// - heap을 이용하자
//  * 장점
//  - 속도 2*N*log2M

// - 큰 값 -> min heap, 작은 값 -> max heap 사용
// - n 개 만큼만 유지한 heap을 가지면 됨

// [-1, 3, -1, 5, 4] 에서 2번째로 큰 값을 찾는 경우
// sol)
// => [-1]
// => [-1, 3]
// => [-1, 3, -1] -> [-1, 3] -> [-1, 3]
// => [-1,3,5] -> [5, 3] -> [3, 5]
// => [3, 5, 4] -> [4, 5] -> [4, 5]

package main

import (
	"dataStruct"
	"fmt"
)

func pickNthValueTest() {
	// heap 을 이용한 nth 값 가져오기
	h := &dataStruct.MinHeap{}

	nums := []int{-1, 3, -1, 5, 4}

	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
		if h.Count() > 2 {
			h.Pop()
		}
	}

	fmt.Println(h.Pop())

	// Input: [2,4,-2,-3,8], 1
	// Output: 8

	h = &dataStruct.MinHeap{}

	nums = []int{2, 4, -2, -3, 8}

	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
		if h.Count() > 1 {
			h.Pop()
		}
	}

	fmt.Println(h.Pop())
}
