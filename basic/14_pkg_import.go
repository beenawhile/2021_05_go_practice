package main

import (
	"create_pkg/testlib"
	"fmt"
)

func importTest() {
	// 외부 패키지 만든 것을 잘 받아오는지 확인해보자
	song := testlib.GetMusic("Alicia Keys")
	fmt.Println(song)
}
