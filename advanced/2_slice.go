// slice : go lang 에서 지원하는 동적 배열
// 동적 배열 원리
// - 동적 배열은 메모리 공간을 포인트하고 있음
// - 동적 배열에 3개를 할당하면 메모리 (쉽게 이야기해서) 3개의 칸을 차지하고 있을 것임
// - 동적 배열에 새로 6개를 할당한다고 하면 메모리 공간 6개를 새로 확보한 후, 이전에 있던 3개 공간의 데이터를 6개로 옮김
// - 동적배열의 포인터를 6개 공간으로 변경
// - 3개 공간을 삭제

// - 매번 공간을 새로 확보한 후 데이터를 이동시키는 것은 비효율적이어, 원래 공간을 확보할 때 여유분을 미리 확보해놓음(이전의 2배)

// 선언 방법
// 1. var a []int
// 2. a := []int{1,2,3,4}
// 3. a := make([]int, 3) // 길이
// 3. a := make([]int, 3,6) // 길이, capacity
package main

import "fmt"

func sliceTest() {

	var a []int
	fmt.Printf("len(a) = %d\ncap(a) = %d\n", len(a), cap(a))
	b := []int{1, 2, 3, 4, 5}
	fmt.Printf("len(b) = %d\ncap(b) = %d\n", len(b), cap(b))
	c := make([]int, 0, 8)
	fmt.Printf("len(c) = %d\ncap(c) = %d\n", len(c), cap(c))

	// 값 추가하는 방법 : append
	a = append(a, 1, 2, 3) // append가 대표적인 Variadic Function(가변인자함수, ...type 사용하는 방법)
	fmt.Println(a)
	fmt.Printf("len(a) = %d\ncap(a) = %d\n", len(a), cap(a))

	// append의 반환값으로 slice가 나옴 => append의 결과로 다른 slice 가 나온다는 의미
	// => 엄밀히 따지면 입력과 출력된 slice가 다를 수 있음
	// capacity 공간이 남아 있을 때는 같은 메모리 공간에 넣지만, capacity 공간이 없을 때는 새로 공간 확보한 후 넣는다 해석할 수 있음
	d := []int{1, 2}
	e := append(d, 3)
	fmt.Printf("d 주소 = %p\ne 주소 = %p\n", d, e)

	// 정밀하게 확인
	for i := 0; i < len(d); i++ {
		fmt.Printf("%d, ", d[i])
	}
	fmt.Println()

	for i := 0; i < len(e); i++ {
		fmt.Printf("%d, ", e[i])
	}
	fmt.Println()

	fmt.Printf("cap(d): %d\ncap(e): %d\n", cap(d), cap(e))

	// 다른경우 확인
	f := make([]int, 2, 4)
	g := append(f, 3)

	fmt.Printf("f주소 = %p\ng주소 = %p\n", f, g) // 주소가 같음
	// cap가 남아서 새로 공간할당하지 않고 그곳에 추가함

	// 같은 메모리 공간을 참고하기 때문에 볼 수 있는 경우
	h := make([]int, 2, 4)
	h[0] = 1
	h[1] = 2

	i := append(h, 3)

	fmt.Printf("h주소 = %p\ni주소 = %p\n", h, i)

	fmt.Println(h) // [1 2]
	fmt.Println(i) // [1 2 3]

	i[0] = 4
	i[1] = 5

	fmt.Println(h) // [4 5]
	fmt.Println(i) // [4 5 3]
	// 둘 다 값이 바뀜 => 같은 주소 보고 있기 때문

	// 이런경우로 에러가 많이 발생하기 때문에 공간을 다르게 쓰고 싶은 경우는 copy 해서 사용하는 것이 바람직함
	j := make([]int, len(h))
	for i := 0; i < len(h); i++ {
		j[i] = h[i]
	}
	j[1] = 63

	fmt.Printf("%p\n", h)
	fmt.Println(h)
	fmt.Printf("%p\n", j)
	fmt.Println(j)

}
