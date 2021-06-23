package main

import (
	"fmt"
)

// golang 에는 for 문 하나만 있음
func forTest() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 0

	for i < 10 {
		fmt.Println(i)
		i++
	}

	fmt.Printf("i = %d\n", i)

	for {
		i++

		if i == 11 {
			continue
		}

		if i == 13 {
			break
		}
		fmt.Println(i)
	}

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5-i; j++ {
			if i < j {
				continue
			}
			fmt.Print("*")
		}

		fmt.Println()
	}

}
