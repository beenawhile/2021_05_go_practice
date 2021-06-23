package main

import "fmt"

func pointerTest() {
	m1, m2 := 2, 3
	ptr := &m1
	fmt.Println(ptr)
	fmt.Println(*ptr)

	// pointer 예시
	fmt.Println(m1, m2)
	fmt.Println(&m1, &m2)
	swap(&m1, &m2)
	fmt.Println(m1, m2)
	fmt.Println(&m1, &m2)

}

// pointer 예시
func swap(m1, m2 *int) {
	temp := *m2
	*m2 = *m1
	*m1 = temp
}
