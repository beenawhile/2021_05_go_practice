package main

import "fmt"

// 대부분의 현대 언어는 pointer 를 감추고 있음
// golang 에서는 pointer의 존재를 명시적으로 꺼내놓되, C 나 C++에서 문제가 되었던 연산이나 casting을 막음
// 심플하고 명확하게 사용하도록 노력한 것으로 보임

func pTest() {
	// pointer는 메모리 주소 포함하고 있음
	// &일반변수 : 일반변수의 주소

	var a int
	var p *int
	p = &a
	a = 3

	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(*p)

	Increase(&a)

	fmt.Println(a)
}

// 포인터로 복사할 때는 적은 양이 메모리에 복사되기 때문에 메모리 상, 성능 상 더 좋음
func Increase(num *int) {
	*num++
}
