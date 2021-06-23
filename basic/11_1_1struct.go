package main

import "fmt"

// tucker's go programming 을 보고 공부한 자료 : struct

// 프로그래밍 언어는 응집성을 높이고, 종속성은 낮추는 방식으로 발전했음

// structure : 응집성을 높이는 방법
//  - 어떠한 특성을 나타내는 정보들이 각각의 field 로 정의되어 있다면 관리하기 까다로울 것임
//  - 하나로 묶어서 하나의 변수로 관리할 수 있게 만들어주는 것

// 기존 C 언어에서는 속성만 있었지만 현대 언어로 오면서 기능이 추가됨
// => first class 라고 부름

type Person2 struct {
	name string
	age  int
}

// method
func (person Person2) GetName() string {
	return person.name
}

type Student struct {
	name  string
	class int
	score Score
}

type Score struct {
	name  string
	grade string
}

func (s *Student) GetScore() {
	fmt.Println(s.score)
}

func (s *Student) InputScore(name, score string) { // {
	s.score.name = name
	s.score.grade = score
}

func structTest() {
	var s Student
	s.name = "철수"
	s.class = 1
	s.score.name = "수학"
	s.score.grade = "C"
	s.GetScore()

	// s.InputScore("과학", "A")
	s.InputScore("과학", "A")
	s.GetScore() // {수학 C} 라고 결과 나옴
	// 중요한 문제
	// golang에서 함수호출의 변수는 무조건 복사로 일어남
	// 함수의 결과가 메모리로 전달되는게 아님
	//  - 함수의 결과를 받아와서 다시 할당하는것 보기
	// 혹은 포인터 !!!
}
