// interface : 두 객체간에 어떤 관계를 가지는지 정의한 것
// - object에서 기능을 외부로 빼냈다고 생각해도 된다함

package main

import "fmt"

type Car2 interface {
	// 다른 클래스가 따르는 function을 정의할 수 있음
	Drive()
	Stop()
}

type Lambo struct {
	LamboModel string
}

type Chevy struct {
	ChevyModel string
}

func (l *Lambo) Drive() {
	fmt.Println("Lambo on the move")
	fmt.Println(l.LamboModel)
}

func (c *Chevy) Drive() {
	fmt.Println("Chevy on the move")
	fmt.Println(c.ChevyModel)
}
func (l *Lambo) Stop() {
	fmt.Println("Stopping lambo")
}

func interfaceTest() {
	l := Lambo{"Gallardo"}
	c := Chevy{"C369"}
	l.Drive()
	c.Drive()
}

// 어떠한 기능을 강제할 수도 있음
// func NewModel(arg string) Car {
// 	return &Lambo{arg}
// }
