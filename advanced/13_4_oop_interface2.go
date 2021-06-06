package main

import (
	"fmt"
	"strconv"
)

type StructA struct {
}

func (a *StructA) AAA(x int) int {
	return x * x
}

func (a *StructA) BBB(x int) string {
	return "X = " + strconv.Itoa(x)
}

type StructB struct {
}

func (b StructB) AAA(x int) int {
	return x * 2
}

type StructC struct {
	val string
}

// string이라는 메소드가 있으면
func (s *StructC) String() string {
	return s.val
}

func interface2Test() {
	var c interfaceA
	c = &StructA{}

	// c = StructA{} // pointer type이 아니어 에러 발생

	// var d interfaceA
	// d = &StructB{} // duck type 법칙에 의해 interfaceA를 모두 구현하지 않았기 때문에 에러 발생

	fmt.Println(c.BBB(3))

	d := &StructC{val: "AAA"}
	fmt.Println(d)
}
