package main

import "fmt"

// OOP의 핵심은 상태와 기능

// OOP로 넘어오면서 만든 object 간의 관계가 가장 중요하게 됨

type BreadOOP struct {
	val string
}

type Jam struct {
}

func (b *BreadOOP) PutJam(jam *Jam) {
	b.val += jam.GetVal()
}

func (j *Jam) GetVal() string {
	return " + jam"
}

func (b *BreadOOP) String() string {
	return b.val
}

func oopTest() {
	bread := &BreadOOP{val: "Bread"}
	jam := &Jam{}

	bread.PutJam(jam)

	fmt.Println(bread)

}
