package main

import (
	"fmt"
)

// arrays : list of ordered items
// - 데이터 타입 정의해줘야함

func arraysTest() {

	// 크기 정해진 배열
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 23
	fmt.Println(a, a[2])

	// 배열 초기화
	var a1 = [3]int{1, 2, 3}
	var a3 = [...]int{1, 2, 3, 4, 5, 6}

	// 배열은 slice 랑 달라서 append 되어지지 않음
	// - 즉, 한번 크기 정해지고 나면 동적으로 크기 변화시킬 수 없음
	// append(a3,23) // 불가능

	fmt.Println(a1, a3)

	// 다차원 배열
	var multiArr [2][3]int
	multiArr[0][0] = 1
	fmt.Println(multiArr)

	var multiArray = [2][2]int{{1, 1}, {1, 1}}
	fmt.Println(multiArray)

}
