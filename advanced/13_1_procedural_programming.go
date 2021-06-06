// OOP 가 아닌 것은 좋은 프로그래밍이 아니다? => OOP는 기존의 문제를 해결하기 위한 하나의 개념, 방법이기 때문에 득과 실이 있다
// OOP를 사용하면서 새로운 문제가 생기게 되고, 그것을 해결하기 위한 방법이 계속 나오고 있다
// OOP는 패러다임일 뿐이고 진화과정일 뿐이지 절대적인 것이 아니다!

// OOP가 나오게 한 문제는 무엇이었나?
// 1. OOP 이전의 방법 : 절차적 프로그래밍(Procedural Programming)
// - 순서대로 따라가면 되기 때문에 이해하기 쉽다
// - 단점 : 비슷한 기능을 하는 다른 기능을 만들고 싶을 때 새로 만들어줘야함 => 유지보수가 힘듬
//          유지보수가 힘들어 지면서 가독성도 떨어지게 됨
//          산탄총 수정을 해야해서 스파게티 코드가 만들어짐
// ※ OOP를 썼다고 코드를 어떻게 짜느냐에 따라 산탄총 수정을 해야할 수도 있음

// Object : 상태 + 기능 => 상태를 어떻게 조정하는가

package main

import "fmt"

// 절차적 프로그래밍
func ppTest() {
	// 1. 빵 두개를 꺼낸다
	breads := GetBreads(2)
	// 2. 딸기잼 뚜껑을 연다
	jam := &StrawberryJam{}
	OpenStrawberryJam(jam)
	// 3. 딸기잼 한스푼 뜬다
	spoon := GetOneSpoon(jam)
	// 4. 딸기잼을 빵에 바른다
	PutJamOnBread(breads[0], spoon)
	// 5. 빵을 덮는다
	sandwitch := MakeSandwitch(breads)
	// 6. 완성
	fmt.Println(sandwitch.val)
}

// 만약 오렌지 잼 샌드위치를 만들고 싶으면? => 기존 함수를 쓸 수 없어서 다시 만들어줘야함

type Bread struct {
	val string
}

type StrawberryJam struct {
	opened bool
}

type SpoonOfStrawberry struct {
}

type SandWitch struct {
	val string
}

func GetBreads(num int) []*Bread {
	breads := make([]*Bread, num)
	for i := 0; i < len(breads); i++ {
		breads[i] = &Bread{val: "bread"}
	}
	return breads
}

func OpenStrawberryJam(jam *StrawberryJam) {
	jam.opened = true
}

func GetOneSpoon(_ *StrawberryJam) *SpoonOfStrawberry {
	return &SpoonOfStrawberry{}
}

func PutJamOnBread(bread *Bread, jam *SpoonOfStrawberry) {
	bread.val += "+ Strawberry Jam"
}

func MakeSandwitch(breads []*Bread) *SandWitch {
	sandwitch := &SandWitch{}
	for i := 0; i < len(breads); i++ {
		sandwitch.val += breads[i].val + " + "
	}
	return sandwitch
}
