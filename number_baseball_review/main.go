package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Results struct {
	strikes int
	balls   int
}

func MakeNumbers() [3]int {
	var nums [3]int
	for i := 0; i < 3; i++ {
		var no int
		var duplicated bool
		for {
			no = rand.Intn(10)
			duplicated = false

			for j := 0; j < i; j++ {
				if nums[j] == no {
					duplicated = true
					break
				}
			}

			if !duplicated {
				nums[i] = no
				break
			}

		}
	}

	fmt.Println(nums)
	return nums
}

func InputNumbers() [3]int {
	var nums [3]int

	var no int

	var success bool

	for {

		fmt.Println("0 ~ 9 까지 중복되지 않는 숫자 3개를 입력하세요.")
		_, err := fmt.Scanf("%d\n", &no)

		if err != nil {
			fmt.Println("잘못 입력하였습니다.")
			continue
		}

		success = true

		idx := 0
		for no > 0 {
			n := no % 10

			no /= 10

			// 중복 확인
			for i := 0; i < idx; i++ {
				if n == nums[i] {
					fmt.Println("중복 번호를 입력하였습니다.")
					success = false
					break
				}
			}

			if idx >= 3 {
				fmt.Println("3자리 이상 입력하였습니다.")
				success = false
				break
			}

			nums[idx] = n
			idx++

		}

		if success && idx < 3 {
			fmt.Println("3개의 숫자를 입력하세요.")
			success = false
		}

		if !success {
			continue
		}

		break

	}

	nums[0], nums[2] = nums[2], nums[0]

	fmt.Println(nums)
	return nums
}

func CompareNumbers(numbers, inputNumbers [3]int) Results {
	var strikes int
	var balls int

	for i, value := range numbers {
		for j, inputValue := range inputNumbers {
			if value == inputValue {
				if i == j {
					strikes++
				} else {
					balls++
				}
			}
		}
	}

	return Results{strikes: strikes, balls: balls}
}

func PrintResults(results Results) {
	fmt.Printf("%dS %dB\n", results.strikes, results.balls)
}

func IsGameEnd(results Results) bool {
	return results.strikes == 3
}

func main() {

	cnt := 0

	rand.Seed(time.Now().UnixNano())

	numbers := MakeNumbers()

	for {
		cnt++

		inputNumbers := InputNumbers()

		results := CompareNumbers(numbers, inputNumbers)

		PrintResults(results)

		if IsGameEnd(results) {
			break
		}
	}

	// 게임 종료. 몇 번만에 맞췄는지 출력
	fmt.Printf("%d번 만에 성공", cnt)

}
