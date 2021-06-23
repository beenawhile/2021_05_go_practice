package main

import "fmt"

// 밖에서 정의할 때는 var로 정의해야함
var globalA = 5

func variableTest() {

	// how to instantiate variable

	// 1
	var a string = "goorm"
	fmt.Println(a)

	var b int = 23
	fmt.Println(b)

	// 2
	var c = true
	fmt.Println(c)

	// 3
	d := "short"
	fmt.Println(d)

	// 4
	var e int
	fmt.Println(e) // 0

	fmt.Println(globalA)

	// how to instantiate multiple variables
	var ab, bc int = 1, 2
	fmt.Println(ab, bc)

	i, j, k := 1, 2, 3
	fmt.Println(i, j, k)

	var str1, str2 string = "Hello", "goorm"
	fmt.Println(str1, str2)

	constTest()

}
