package main

import (
	"fmt"
	"math"
)

// interface
// struct가 필드들의 집합체라면, interface는 메소드 들의 집합체
// - 타입이 구현해야 하는 메소드 원형들을 정의
// - 하나의 사용자 정의 타입이 interface를 구현하기 위해서는 단순히 그 인터페이스가 갖는 모든 메소드들을 구현하면 됨

// interface 정의
type Shape interface {
	area() float64
	perimeter() float64
}

// interface 구현
//Rect 정의
type Rectangle struct {
	width, height float64
}

//Circle 정의
type Circle struct {
	radius float64
}

//Rect 타입에 대한 Shape 인터페이스 구현
func (r Rectangle) area() float64 { return r.width * r.height }
func (r Rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

//Circle 타입에 대한 Shape 인터페이스 구현
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// interface 사용
func interfaceNewTest() {
	r := Rectangle{10., 20.}
	c := Circle{10}

	showArea(r, c)

	// interface type (=empty interface) : 다양한 값이 들어갈 수 있을 때 사용하는 type
	// - interface{} (=empty interface) 로 사용
	// - empty interface : 메서드를 전혀 갖지 않는 빈 인터페이스, 모든 타입을 나타내기 위해 빈 인터페이스 사용함
	// - = dynamic type

	var x interface{}
	x = 1
	fmt.Println(x)
	x = "string"
	fmt.Println(x)
	x = struct{}{}
	fmt.Println(x)

	// type assertion
	// - interface type의 x와 타입 T에 대하여 x.(T) 로 표현했을 때,
	//   이는 nil이 아니며, x는 T 타입에 속한다는 점을 확인(assert) 하는 것
	// - nil 이거나 타입이 T가 아닐경우 runtime error 발생

	// var a interface{} // 이런식으로 정의하지 않았을 경우 fatal panic 을 내보냄
	var a interface{} = 1

	i := a
	j := a.(int)

	fmt.Println(i)
	fmt.Println(j)

}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area()
		fmt.Println(a)
	}
}

// interface type (=empty interface) : 다양한 값이 들어갈 수 있을 때 사용하는 type
// - interface{} (=empty interface) 로 사용
// - empty interface : 메서드를 전혀 갖지 않는 빈 인터페이스, 모든 타입을 나타내기 위해 빈 인터페이스 사용함
// - = dynamic type
