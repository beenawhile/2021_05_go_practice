package main

import "fmt"

// function과 method는 go에서 다른의미로 적용되니 잘 이해하기

func funcTest() {
	x := add(32, 24)
	fmt.Println(x)

	var (
		y = 3
		z = 4
	)

	yy, zz := return2(y, z)
	fmt.Printf("두개 인자 넣어서 두개 호출하는 것도 가능 \n %d, %d 두개 인자 넣었을 때 제곱값: %d %d", y, z, yy, zz)

	saySomthing("This", "is", "numerous", "variables")
	saySomthing("test")

	// Anonymous Function(익명함수)
	sum := func(n ...int) int { //익명함수 정의
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}

	result := sum(1, 2, 3, 4, 4, 3, 2, 1)
	println(result)

	// return 을 특이하게 정의하는 방식 => 정의가서 보기
	helloGo := sayHello()
	println(helloGo)

	// 일급함수 : 정의가서 확인
	r1 := calc(func(x int, y int) int { return x - y }, 10, 5)
	println(r1)

	// 함수원형 사용
	r2 := calc(func(x int, y int) int { return x * y }, 10, 5)
	println(r2)

	// closure 사용
	next := nextValue()
	println(next)
	println("-----------------")

	println(next())
	println(next())
	println(next())
	println("-----------------")

	anotherNext := nextValue()
	println(next())
	println(next())
	println(next())
	println("-----------------")
	println(anotherNext())
	println(anotherNext())
	println(anotherNext())

}

func add(x int, y int) int {
	return x + y
}

func return2(x, y int) (int, int) {
	return x * x, y * y
}

// Pass By Value
// - 값을 복사하기 때문에 함수 내에서 값이 변경된다 하더라도 바뀌지 않음
func say(msg string) {
	println(msg)
	msg = "Changed" //메시지 변경
}

// Pass By Reference
// - 변수 앞에 &를 포인터 붙여 사용
// - 주소 값을 넣어 값을 변경하면 변경됨
func sayAndChange(msg *string) {
	println(*msg)
	*msg = "Changed" //메시지 변경
}

// Variadic Function (가변인자함수)
// - 고정된 수의 파라미터들을 전달하지 않고 다양한 갯수의 파라미터를 전달하고자 할때
// 변수 앞에 ... 사용
func saySomthing(msg ...string) {
	for _, s := range msg {
		println(s)
	}
}

// return 을 다음 처럼 해줄 수도 있음
func sayHello() (helloGo string) {
	helloGo = "Hello, Go"
	return
}

// 일급함수 : Go 에서는 함수가 기본 타입과 동일하게 처리 되기 때문에 파라미터로 전달하거나 리턴값으로 사용할 수 있음
func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

// type문 사용하여 함수 원형 정의
// - type문은 구조체, 인터페이스 등 custom type을 정의하기 위해 사용
// - 함수 원형 정의하는데도 사용
// - 반복되서 사용하는 함수 원형을 정의
type calculator func(a int, b int) int

// - 이용
func calc2(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}

// closure : 함수 바깥에 있는 변수를 참조하는 함수값
// - 이 때 함수는 바깥의 변수를 마치 함수 안으로 끌어들인 듯이 그 변수를 읽거나 쓸 수 있게 됨
// - go 에서는 함수를 closure 로서 사용할 수 있음
func nextValue() func() int {
	i := 0
	// 이 i 값은 반환하는 익명함수 바깥에 있지만 안에 있는 것 처럼 끌어와 사용할 수 있음
	// 예제 참고
	return func() int {
		i++
		return i
	}
}
