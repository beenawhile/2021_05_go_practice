// map = dictionary
//  - key-value 타입으로 이루어진 자료구조
//  - 단순히 늘어놓기만 하면 read 시에 O(N) 의 속도
//  - BST(Binary Search Tree)를 이용하여 속도를 높이는 방법을 사용함 => sorted map, ordered map
//  	- 속도 : log2N

// hash map
//  - 같은 입력 -> 같은 출력
//  - 다른 입력 -> 다른 출력
//  - 배열에 값을 만들고, hash 라는 함수에 key 값을 넣어 나온 인덱스 값을 이용해 배열에서 해당 인덱스 값을 불러옴
//  - 속도 : O(1) => 일정한 속도 => 시간이 거의 걸리지 않음

// * 관건
//  - hash 를 어떻게 만드는가?
//    1. 출력값의 범위가 정해져 있음
//    2. 같은 입력이면 같은 출력이 나옴
//    3. 다른 입력이면 (보통의 경우) 다른 출력이 나옴
//   - sine : 출력값 범위 정해져 있음
//   - mod : 출력값 범위 정해져 있음, 연산 단순함 => hash에 많이 사용된다 함
//     - one-way function : 나눔값과 나머지를 통해 원래 값을 추론할 수 없음(추론값이 너무 많기 때문) => 한 방향으로만 계산할 수 있기 때문에 one way function
//     - 즉, 암호문을 만들 수 있다는 것임

package main

import (
	"dataStruct"
	"fmt"
)

func customMapTest() {
	fmt.Println("abcde = ", dataStruct.Hash("abcde"))
	fmt.Println("abcde = ", dataStruct.Hash("abcde"))
	fmt.Println("abcdef = ", dataStruct.Hash("abcdef"))
	fmt.Println("absdcsdcdswefsdfsdff = ", dataStruct.Hash("absdcsdcdswefsdfsdff"))

	// map
	customMap := dataStruct.CreateMap()
	customMap.Add("AAA", "010777777777")
	customMap.Add("BBB", "010888888888")
	customMap.Add("CCC", "010999999999")

	fmt.Println("AAA = ", customMap.Get("AAA"))
	fmt.Println("CCC = ", customMap.Get("CCC"))
	fmt.Println("DDD = ", customMap.Get("DDD"))
}
