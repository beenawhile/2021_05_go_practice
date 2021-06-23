package main

import (
	"fmt"
)

func mapTest() {
	// 일반적인 map 표현 방식
	// var idMap map[string]interface{}

	// var idMap map[string]string
	// 이때 선언된 변수는 nil 값을 갖음(map은 reference 타입이므로) => nil map
	// nil 맵에는 어떠한 데이터를 쓸 수 없음 => make 사용하여 초기화 해줘야함
	// idMap["test"] = "test"
	// fmt.Println(idMap)

	// 초기화 방법1 : make 함수
	var idMap = make(map[string]string)
	// idMap["test"] = "test"
	fmt.Println(idMap)

	// 초기화 방법2 : literal 사용
	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "Facebook",
	}
	fmt.Println(tickers)

	// map 사용
	var m map[int]string
	m = make(map[int]string)

	m[901] = "Apple"
	m[134] = "Grape"
	m[777] = "Tomato"

	str := m[134]
	fmt.Println(str)

	noData := m[999] // 값이 없으면 nil 혹은 zero 리턴
	fmt.Println(noData)

	// 삭제
	delete(m, 777)
	fmt.Println(m)

	// map 키 값 체크
	tickersNullCheck := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "Facebook",
	}

	val, exists := tickersNullCheck["TSLA"]
	if !exists {
		fmt.Println("No such key")
	} else {
		fmt.Println(val)
	}

	// for loop 이용한 열거
	for key, val := range tickers {
		fmt.Printf("key: %s, value : %v\n", key, val)
	}

}
