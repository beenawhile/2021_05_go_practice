package main

import "fmt"

func operatorTest() {
	// binary operator
	// 1. 산술연산자(arithmetic operator)
	//  : +, -, *, /, %
	num1, num2 := 17, 5
	// str1, str2 := "Hello", "goorm"
	fmt.Println("num1 % num2 =", num1%num2)

	// 2. 증감연산자(increment and decrement operator)
	count1, count2 := 1, 10.4
	count1++
	count2--
	fmt.Println("count1++ : ", count1)
	fmt.Println("count2-- : ", count2)

	// 3. 할당연산자
	// =,		:=	, +=	, -=	, *=	, /=	, %=	,
	// &= : AND 비트 연산 후 대입	,
	// |= : OR 비트 연산 후 대입	,
	// ^= : XOR 비트 연산 후 대입	,
	// &^= : AND NOT 비트 연산 후 대입	,
	// <<= : 비트를 왼쪽으로 이동 후 대입,
	// >>= : 비트를 오른쪽으로 이동 후 대입

	// print에 수식 바로 못넣으니 참고

	a := 3
	var num int
	num = a
	fmt.Println("num = a :", num)
	num &= 2
	fmt.Println("num &= 2 :", num)
	num |= 5
	fmt.Println("num |= 5 :", num)
	num ^= 4
	fmt.Println("num ^= 4 :", num)
	num &^= 2
	fmt.Println("num &^= 2 :", num)
	num <<= 9
	fmt.Println("num &^= 2 :", num)
	num >>= 8
	fmt.Println("num &^= 2 :", num)

	// 4. 논리연산자(logical operator)
	// - &&, ||, !
	var boolA bool = true
	boolB := false

	fmt.Println("0 && 0 : ", boolB && boolB)
	fmt.Println("0 && 1 : ", boolB && boolA)
	fmt.Println("1 && 1 : ", boolA && boolA)
	fmt.Println("0 || 0 : ", boolB || boolB)
	fmt.Println("0 || 1 : ", boolB || boolA)
	fmt.Println("1 || 1 : ", boolA || boolA)

	fmt.Println("!1 ", !true)
	fmt.Println("!0 ", !false)

	// 5. 관계 연산자(relational operator)
	// 6. 비트 연산자(bitwise operator)
	// &	|	 ^	&^ << >>

	bit1 := 15 //00001111
	bit2 := 20 //00010100

	fmt.Printf("bit1 & bit2 : %08b, %d\n", bit1&bit2, bit1&bit2)
	fmt.Printf("bit1 | bit2 : %08b, %d\n", bit1|bit2, bit1|bit2)
	fmt.Printf("bit1 ^ bit2 : %08b, %d\n", bit1^bit2, bit1^bit2)
	fmt.Printf("bit1 &^ bit2 : %08b, %d\n", bit1&^bit2, bit1&^bit2)

	fmt.Printf("bit1 << 4 : %08b, %d\n", bit1<<4, bit1<<4)
	fmt.Printf("bit2 >> 2 : %08b, %d\n", bit2>>2, bit2>>2)

	// 7. 채널 연산자(channel operator)
	// - 채널 사용할 때 쓰는 연산자
	// - 간단히 말해, channel과 goroutine간 데이터를 주고 받고 실행흐름을 제어하는 기능
	ch := make(chan int) // 정수형 채널 생성

	go func() {
		ch <- 10
	}() // 채널에 10을 보냄

	result := <-ch // 채널로부터 10을 전달 받음
	fmt.Println(result)

	// 8. 포인터 연산자(pointer operator)
	// - & : 변수 메모리 주소를 참조,  * : 포인터 변수에 저장된 메모리에 접근하여 값을 참조
	pointNum1 := 5
	pointNum2 := &pointNum1

	fmt.Println(pointNum2)
	fmt.Println(*pointNum2)

	// go 연산자 우선순위는 표 보면서 참고하기

	// 콘솔 입력 해보기
	var scanNum1, scanNum2, scanNum3 int

	fmt.Print("정수 3개를 입력하세요 : ")
	fmt.Scanln(&scanNum1, &scanNum2, &scanNum3)
	fmt.Println(scanNum1, scanNum2, scanNum3)

}
