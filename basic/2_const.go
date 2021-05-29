package main

import "fmt"

const username = "kim"

func constTest() {
	// 상수는 var 키워드 대신에 const 키워드를 사용하고 생략할 수 없기 때문에 자연스럽게 := 용법을 사용할 수 없음.
	const a int = 1
	const b, d = 20, 10
	const c = "goorm"
	fmt.Println(username)

	// 변수와 다르게 괄호 이용해서 여러 개 값을 묶어서 초기화 할 수 있음
	const (
		// 규칙
		// 1. 괄호 묶여 있는 상수들은 다른형으로 초기화 가능
		// 2. 괄호 시작, 마지막 위치는 상관없지만 각 상수들은 개행해서 초기화
		// 3. 콤마 입력하면 안됨
		// 4. 묶어서 선언된 상수들 중 첫번째 값은 꼭 선언해야함, 선언되지 않은 값은 전 상수 값을 가짐
		// 5. iota 라는 식별자를 값으로 초기화하면 그 후에 초기화하지 않고 이어지는 상수들은 순서(index)가 값으로 저장됨
		c1 = 10
		c2
		c3
		c4
		c5
		c6 = iota
		c7
		c8
		c9 = "earth"
		c10
		c11 = "End"
	)

	fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11)

}
