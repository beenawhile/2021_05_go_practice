package main

import (
	"fmt"
)

// arrays : list of ordered items
// - 데이터 타입 정의해줘야함

func sliceTest() {
	// 슬라이스 생성방법 1
	var slice1 []int
	slice1 = []int{1, 2, 3, 4}
	fmt.Println(slice1)

	slice2 := []int{1, 2, 3, 4}
	slice2[3] = 23
	fmt.Println(slice2)

	// 추가는 무조건 append 로만 가능
	slice2 = append(slice2, 42, 325, 235667)
	fmt.Println(len(slice2), cap(slice2)) // 길이: 7, 용량: 8
	fmt.Println(slice2)

	// 슬라이스 생성방법 2 : make
	// 장점 : length와 capacity을 임의로 지정할 수 있음
	slice3 := make([]int, 5, 10)
	fmt.Println(len(slice3), cap(slice3)) // 길이: 5, 용량: 10
	fmt.Println(slice3)

	slice4 := make([]int, 5)
	fmt.Println(len(slice4), cap(slice4)) // 길이: 5, 용량: 5
	fmt.Println(slice4)

	slice4 = append(slice4, 23, 45)
	fmt.Println(len(slice4), cap(slice4)) // 길이: 7, 용량: 10
	fmt.Println(slice4)

	// array 비교
	var arr = [...]int{1, 2, 3, 4}
	fmt.Println(len(arr), cap(arr))
	fmt.Println(arr)

	// nil slice
	var nilSlice []int
	if nilSlice == nil {
		fmt.Println("Nil Slice")
	}
	fmt.Println(len(nilSlice), cap(nilSlice)) // 모두 0

	// 부분 슬라이스
	slice4 = slice4[:]
	fmt.Println(slice4)

	slice4 = slice4[3:]
	fmt.Println(slice4)

	slice4 = slice4[len(slice4)-2 : len(slice4)-1]
	fmt.Println(slice4)

	// append
	// 슬라이스 용량(capacity)이 아직 남아 있는 경우
	// - 그 용량 내에서 슬라이스의 길이(length)를 변경하여 데이타를 추가하고,
	// 용량(capacity)을 초과하는 경우
	// - 현재 용량의 2배에 해당하는 새로운 Underlying array을 생성하고 기존 배열 값들을 모두 새 배열에 복제한 후 다시 슬라이스를 할당한다.
	// len=0, cap=3 인 슬라이스
	sliceCapCheck := make([]int, 0, 3)

	// 계속 한 요소씩 추가
	for i := 1; i <= 15; i++ {
		sliceCapCheck = append(sliceCapCheck, i)
		// 슬라이스 길이와 용량 확인
		fmt.Println(len(sliceCapCheck), cap(sliceCapCheck))
	}

	fmt.Println(sliceCapCheck) // 1 부터 15 까지 숫자 출력

	// slice 2개 확장하기
	sliceA := []int{1, 2, 3}
	sliceB := []int{4, 5, 6}
	sliceA = append(sliceA, sliceB...)
	fmt.Println("----------")
	fmt.Println(sliceA)

	// copy
	source := []int{1, 2, 3}
	target := make([]int, len(source), cap(source)*2)
	copy(target, source)
	fmt.Println(target)
	fmt.Println(len(target), cap(target))

}
