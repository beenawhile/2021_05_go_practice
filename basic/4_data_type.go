package main

import "fmt"

func dataTypeTest() {

	// 따로 지정하지 않으면 기본적으로 int(n비트 시스템에서 n비트), float64 할당

	// 1. bool : 1byte
	// 2. 정수 : int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr
	// 3. 실수 및 복소수 : (실수) float32, float64, (복소수) complex64, complex128
	// 4. 문자열 타입 : string [16byte], immutable 타입이어 값을 수정할 수 없음
	// 5. 기타타입: (정수 0, 양수) byte == uint8 [1 byte], (정수) rune [4byte]

	// 문자열 표현 방법
	// 1. 백쿼트 : `` => Raw String Literal => 이스케이프 시퀀스가 의미로 인식되지 않고 raw 하게 적용
	// 2. 이중인용부호 : "" => Interpreted String literal => 이스케이프 시퀀스 적용됨
	// - 둘 다 +로 문자열 합칠수 있음

	var rawLiteral = `바로 실행해보면서 배우는 \n Golang`
	var interLiteral = "바로 실행해보면서 배우는 \nGolang"

	plusString := "구름" + "EDU\n" + "GoLang"

	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)
	fmt.Println()
	fmt.Println(plusString)

	// 자료형 변환
	// - go 에서는 명시적으로 형 변환을 지정해줘야함
	var num int = 10
	var chagef float32 = float32(num)
	changei := int8(num)

	var str string = "goorm"
	changestr := []byte(str)
	str2 := string(changestr)

	fmt.Println(num)
	fmt.Println(chagef, changei)

	fmt.Println(str)
	fmt.Println(changestr)
	fmt.Println(str2)

	// 주의할점
	// var num1, num2 int = 3, 4
	// var result float32 = num1 / num2	// 에러 발생 : 자동으로 형변환 하지 않기 때문에 결과로 넣은 값을 float32로 형변환해줘야함
	// fmt.Printf("%f", result)

}
