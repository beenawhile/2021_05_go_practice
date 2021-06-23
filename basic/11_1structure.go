package main

import "fmt"

// structure 구조체 : abstract data type
// - 논리적으로 연관되어 있는 데이터끼리 묶여있음
// - data encapsulation 이라 부름

// struct : Custom Data Type 표현할 때 사용
//  - field 들의 집합체, 필드들의 컨테이너
//  - 필드 데이터만 가지고 메소드는 갖지 않음

// Go는 OOP를 고유의 방식으로 지원 => 전통적인 OOP언어가 가지는 클래스, 객체, 상속의 개념이 없음
// 전통적인 OOP의 class => struct로 표현
// 메소드는 별도로 분리하여 정의

// private은 소문자, Public은 대문자로 정의해야한다는 것 기억

// struct는 기본적으로 mutable 개체
// - 다른 함수로 파라미터를 넘기면 pass by value에 따라 객체를 복사해서 전달

type Car struct {
	Name    string
	Age     int
	ModelNo int
}

type person struct {
	name string
	age  int
}

func structureTest() {
	// how to instantiate
	c := Car{}
	// var c1 Car
	fmt.Println(c) // 출력: {" " 0 0}
	c1 := Car{"chevy", 1, 2123}
	fmt.Println(c1) // {chevy 1 2123}
	// other technique
	c2 := Car{
		Name:    "chevy",
		Age:     1,
		ModelNo: 2315,
	}
	fmt.Println(c2) // {chevy 1 2315}

	c2.Print()
	c2.Drive()
	c2.getName()

	fmt.Println("------------")

	// struct 선언
	// 객체 생성 1
	p := person{}
	p.name = "Shawn"
	p.age = 23
	fmt.Println(p)
	fmt.Println(&p)

	// 객체 생성 2
	var p1 person
	p1 = person{name: "Bob", age: 20}
	p2 := person{name: "Sean", age: 50}

	fmt.Println(p1, p2)
	fmt.Println(&p1, &p2)

	// 객체 생성 3
	p3 := new(person)
	p3.name = "Newman"
	fmt.Println(p3)

	fmt.Println(&p3)

	// 생성자 함수
	// - 사용 전에 초기화되야 하는 경우 좋음
	dic := newDict()
	dic.data[1] = "A"
	fmt.Println(*dic)

}

// extension 기능인듯
func (c Car) Print() {
	fmt.Println(c)
}

func (c Car) Drive() {
	fmt.Println("driving...")
}

func (c Car) getName() string {
	return c.Name
}

type dict struct {
	data map[int]string
}

// 생성자 함수
func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d
}
