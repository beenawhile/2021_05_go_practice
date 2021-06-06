package main

import (
	"fmt"
)

type Key struct {
	k int
}

type Value struct {
	v int
}

func mapTest() {
	// map[key type]value type
	var m map[string]string
	// var m1 map[int]string
	// var m2 map[Key]Value
	// var m3 map[Key]*Value

	// 선언만 해서 쓸수는 없음
	// m["aaa"] = "123"
	// 초기화 시켜줘야함
	m = make(map[string]string)
	m["aaa"] = "aaa"
	fmt.Println(m["aaa"])

	// 없는 경우 빈문자열 나옴
	fmt.Println(m["aaaa"])

	// 선언과 초기화 동시에
	m1 := make(map[int]string)
	m1[543] = "1234"
	fmt.Println(m1[543])

	m2 := make(map[int]int)
	m2[234] = 234
	// 없는 값은 value의 기본형 값이 나옴
	fmt.Println(m2[12]) // 0

	// 이 기본값 성질을 이용
	// C++의 set을 다음과 같이 사용할 수 있음
	// 값이 설정 되어 있는지 아닌지로 확인
	m3 := make(map[int]bool)
	m3[23] = true
	fmt.Println(m3[23], m3[234])

	// 이 기본값 때문에 불편한 점 : 값이 없어서 기본 값이 뜬 것인지, 값을 설정했는데 기본 값인건지 알 수 없음
	m4 := make(map[int]int)
	m4[0] = 123
	fmt.Println("m4[0] = ", m4[0])
	m4[1] = 0
	fmt.Println("m4[1] = ", m4[1])
	fmt.Println("m4[2] = ", m4[2])

	// 이 문제를 타계하기 위해서 golang 에서는 value와 함께 값이 있는지 없는지도 같이 제공
	v, ok := m4[0]
	v1, ok1 := m4[1]
	v2, ok2 := m4[2]
	fmt.Println(v, ok)
	fmt.Println(v1, ok1)
	fmt.Println(v2, ok2)

	// 값 지우기
	// delete(map, key value)
	delete(m4, 1)
	v1, ok1 = m4[1]
	fmt.Println(v1, ok1)

	// 순회
	// - 정렬 안되는 것 주의!
	m4[23] = 2354
	m4[2345] = 213
	m4[234] = 123

	for key, value := range m4 {
		fmt.Println(key, " : ", value)
	}

}
