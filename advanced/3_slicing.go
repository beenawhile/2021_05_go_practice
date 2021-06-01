package main

import "fmt"

func slicingTest() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[4:8]
	fmt.Println(b)
	fmt.Println(a[:5])
	fmt.Println(a[:])
	fmt.Println(a[:8])

	fmt.Printf("%p, %p\n", a, b)

	// 그냥 값을 참고하고 있다는 것 확인해보기
	b[0] = 0
	b[1] = 1
	fmt.Println(a)

	// 뒤에서부터 삭제(실제로는 안보이게 하는)하는 기능은 어떻게 구현해야할까?
	for i := 0; i < 5; i++ {
		var back int
		a, back = RemoveBack(a)
		// 맨 뒷 값을 반환하는 경우는?
		fmt.Println(back)
	}

	fmt.Println(a)

	// 앞에서부터 삭제(실제로는 안보이게 하는)하는 기능은?
	for i := 0; i < 3; i++ {
		var front int
		a, front = RemoveFront(a)
		fmt.Println(front)
	}

	fmt.Println(a)

	// 실제로 메모리 상에 안보이는 값들은 남아있음
	fmt.Printf("%d\n%d\n", len(a), cap(a))

	// 복사만 하고 싶을 때는 새로운 slice 하나 만들고 복사해야함
	c := make([]int, len(a))
	for i := 0; i < len(c); i++ {
		c[i] = a[i]
	}
}

func RemoveBack(a []int) ([]int, int) {
	return a[:len(a)-1], a[len(a)-1]
}

func RemoveFront(a []int) ([]int, int) {
	return a[1:], a[0]
}
