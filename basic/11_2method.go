package main

import "fmt"

// 구조체에서 사용할 수 있는 메소드는 별도로 분리되서 정의됨
// - 특별한 형태의 func 함수
// - func 키워드와 함수명 사이에 어떤 struct를 위한 메소드인지 작성해야함 => receiver 라고 불리는 부분

type Rect struct {
	width, height int
}

func (r Rect) area() int {
	return r.width * r.height
}

func methodTest() {
	rect := Rect{width: 20, height: 20}
	area1 := rect.area()
	fmt.Println(rect.width, area1)

	// pointer receiver 예시
	area := rect.area2()
	fmt.Println(rect.width, area)

	// value receiver와 비교
	area2 := rect.area3()
	fmt.Println(rect.width, area2)
}

// Value receiver vs Pointer receiver
// - value receiver : struct의 데이터를 복사하여 전달
//  * 메서드내에서 그 struct의 필드값이 변경되더라도 호출자의 데이터는 변경되지 않음
// - pointer receiver : struct의 포인터만 전달
//  * 메서드내에서 그 struct의 필드값이 변경이 그대로 호출자의 데이터에 반영

// pointer receiver 예시
func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

func (r Rect) area3() int {
	r.width++
	return r.width * r.height
}
