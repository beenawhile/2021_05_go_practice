package main

import (
	"fmt"
)

// arrays : list of ordered items
// - 데이터 타입 정의해줘야함

func arraysTest() {

	// 크기 정해진 배열
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 23
	fmt.Println(a, a[2])

	// 배열 초기화
	var a1 = [3]int{1, 2, 3}
	var a3 = [...]int{1, 2, 3, 4, 5, 6}

	// 배열은 slice 랑 달라서 append 되어지지 않음
	// - 즉, 한번 크기 정해지고 나면 동적으로 크기 변화시킬 수 없음
	// append(a3,23) // 불가능

	fmt.Println(a1, a3)

	// 다차원 배열
	var multiArr [2][3]int
	multiArr[0][0] = 1
	fmt.Println(multiArr)

	var multiArray = [2][2]int{{1, 1}, {1, 1}}
	fmt.Println(multiArray)

	var testArr [10]int
	for i := 0; i < len(testArr); i++ {
		testArr[i] = i * i
	}

	fmt.Print(testArr)
	fmt.Println()

	// 문자열도 배열
	// ex) s:="Hello, world" 는 총 11byte(11개 문자)로 이루어진 배열
	// byte = uint8
	// s == [11]byte

	s := "Hello world"

	for i := 0; i < len(s); i++ {
		// fmt.Print(s[i], ", ")
		fmt.Print(string(s[i]), ", ")
	}
	fmt.Println()

	// 한글은?

	sKo := "헬로우 월드"
	// 길이는? 16 => 각 한글 글자가 3바이트씩이라는 것을 알 수 있음
	fmt.Println("len(sKo) : ", len(sKo))
	for i := 0; i < len(sKo); i++ {
		// fmt.Print(string(sKo[i]), ", ") // í, , ¬, ë, ¡, , ì, , °,  , ì, , , ë, , ,
		fmt.Print(sKo[i], ", ")
	}
	fmt.Println()

	// 한글은 이상하게 나옴
	// - 한글은 utf-8(1~3byte)이기 때문에 1byte 단위로 끊어서 적으면 이상하게 나오게 됨

	// 한글을 한글자씩 찍어보여주는 방법은? => []byte -> []rune 배열로 바꿔주면 됨
	// rune : UTF-8 문자를 나타내는 타입 (1~3 byte)

	sKoRune := []rune(sKo)
	fmt.Println("len(sKoRune) : ", len(sKoRune))
	for i := 0; i < len(sKoRune); i++ {
		// fmt.Print(sKoRune[i], ", ") // 54764, 47196, 50864, 32, 50900, 46300,
		fmt.Print(string(sKoRune[i]), ", ")
	}
	fmt.Println()

	// 배열복사
	arr := [5]int{1, 2, 3, 4, 5}
	clone := [5]int{}

	for i := 0; i < len(arr); i++ {
		clone[i] = arr[i]
	}
	fmt.Println("clone:", clone)
	fmt.Println("arr:", arr)

	// 배열 역순 만들기?
	arr = [5]int{1, 2, 3, 4, 5}
	temp := [5]int{}
	for i := 0; i < len(arr); i++ {
		temp[i] = arr[len(arr)-1-i]
	}

	arr = temp

	fmt.Println("arr : ", arr)
	fmt.Println("temp : ", temp)

	// 복사하지 않고 자리만 바꾸는 방법
	arr = [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr)/2; i++ {
		// 이중 대입
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}

	fmt.Println("arr : ", arr)

	// 정렬 알고리즘 중 1개 : RADIX 알고리즘
	//  - 가장 빠르면서 간단한 알고리즘
	//  - 단점 : 모든 경우에 사용할 수 없고, 특정한 경우에만 사용할 수 있음
	//  - 특정 범위 숫자를 정해놓고 리스트를 돌면서 각 숫자에 몇개씩 있는지 센 다음, 센 결과를 바탕으로 리스트 작성
	//  - 범위 숫자가 정해져있어야 가능
	//  - 정수가 아닌 실수일 경우 불가능
	sortArr := [11]int{0, 6, 4, 6, 8, 3, 4, 6, 7, 4, 9}
	tempArr := [10]int{}
	for i := 0; i < len(sortArr); i++ {
		idx := sortArr[i]
		tempArr[idx]++
	}

	idx := 0

	for i := 0; i < len(tempArr); i++ {
		for j := 0; j < tempArr[i]; j++ {
			sortArr[idx] = i
			idx++
		}
	}

	fmt.Println("sortArr : ", sortArr)

	// 사용할 수 있는 조건 예시
	// 1. 이름을 앞글자만 가지고 정렬하고 싶을 때, 알파벳 26글자만 있으면 되니(=값의 범위가 한정적) 앞에서 하나씩 뽑아서 정렬하면 됨
	// 알고리즘 속도 : N개

}
