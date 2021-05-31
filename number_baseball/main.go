package main

import "fmt"

func MakeNumbers() [3]int{
	var rst [3]int
	return rst
}

func InputNumbers() [3]int{
	var rst [3]int
	return rst
}

func CompareNumbers(numbers, inputNumbers [3]int) bool{
	var rst bool
	return rst
}

func PrintResult(result bool){
	fmt.Println(result)
}

func IsGameEnd(results bool) bool{
	return true
}

func main() {
	cnt := 0
	for {
		cnt++
		// 무작위 숫자 3개 만듬
		numbers := MakeNumbers()
		// 사용자 입력을 받음
		inputNumbers := InputNumbers()
		// 결과를 비교
		results := CompareNumbers(numbers, inputNumbers)
		// 결과 출력
		PrintResult(results)
		// 3S 인가?
		if IsGameEnd(results) {
			// 게임 끝
			break
		}
	}

	// 게임 종료. 몇 번만에 맞췄는지 출력
	fmt.Printf("%d번 만에 성공", cnt)
}
