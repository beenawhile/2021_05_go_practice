package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ifElse() {
	fmt.Println("숫자를 입력하세요.")
	reader := bufio.NewReader(os.Stdin) // 값 입력 받기
	line, _ := reader.ReadString('\n')  // \n이 나올때 까지 읽기 (single quote 인 것 보기), 결과, 에러로 2개 결과 나옴
	line = strings.TrimSpace(line)      // 빈 공간 날리기
	// 문자열 -> 숫자열로 변경
	n1, _ := strconv.Atoi(line)

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	n2, _ := strconv.Atoi(line)

	fmt.Printf("입력하신 숫자는 %d, %d 입니다.", n1, n2)

	fmt.Println("연산자를 입력하세요")

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	if line == "+" {
		fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
	} else if line == "-" {
		fmt.Printf("%d - %d = %d", n1, n2, n1-n2)
	} else if line == "/" {
		fmt.Printf("%d / %d = %d", n1, n2, n1/n2)
	} else if line == "*" {
		fmt.Printf("%d * %d = %d", n1, n2, n1*n2)
	} else {
		fmt.Println("잘못된 연산자를 입력했습니다.")
	}

}
