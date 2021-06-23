package main

import "fmt"

// interface : 객체간 상호관계를 따로 정의한 것
// - OOP = status + function
//   - function = internal + external method
//     - internal method : 내부에서의 기능을 표현
//     - external method : 외부와의 관계를 표현

// - external method를 따로 정의한 것을 interface라고 함
// - * 기능부분만 decoupling한 것을 interface *
//   decoupling : 종속성을 때냈다, 의존관계를 없앴다

// GoLang의 OOP가 다른 언어의 OOP와 다른점
//  - Duck Typing : 오리라고 꼬리표를 달지 않아도 오리처럼 소리내고 오리처럼 걷고 오리처럼 헤엄치면 그건 오리다
//    - implement 했다고 표시하지 않아도 같은 메소드를 구현하고 있으면 부모를 상속한 것이다
//  - python 에서도 지원함

// 현대 프로그래밍의 발전 방향
// - 종속성 ↓, 응집성 ↑
// - interface는 이를 하기 위한 도구
// - 응집성은 어떠한 객체를 기준으로 관계있는 상태와 기능이 묶임
// - 종속성을 낮추면 확장성이 좋아짐
// - 관계 있는 것은 묶어주고 독립시켜야할 것은 독립시킴

type IJam interface {
	GetOneSpoon() SpoonOfJam
}

type SpoonOfJam interface {
	String() string
}

type BreadObject struct {
	val string
}

type StrawberryJamObject struct {
}

type SpoonOfStrawberryJam struct {
}

type OrangeJam struct{}

type SpoonOfOrangeJam struct {
}

func (s *SpoonOfOrangeJam) String() string {
	return "+ Orange"
}

func (j *OrangeJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{}
}

func (s *SpoonOfStrawberryJam) String() string {
	return " + Strawberry"
}

func (j *StrawberryJamObject) GetOneSpoon() SpoonOfJam {
	return &SpoonOfStrawberryJam{}
}

func (b *BreadObject) PutJam(jam IJam) {
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}

func (b *BreadObject) String() string {
	return "Bread " + b.val
}

// 새 잼 추가

type AppleJam struct {
}

func (j *AppleJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfAppleJam{}
}

type SpoonOfAppleJam struct {
}

func (s *SpoonOfAppleJam) String() string {
	return "+ Apple"
}

func oopInterfaceTest() {
	bread := &BreadObject{}
	// jam := &StrawberryJamObject{}
	// jam := &OrangeJam{}
	jam := &AppleJam{}
	bread.PutJam(jam)

	fmt.Println(bread)

}
