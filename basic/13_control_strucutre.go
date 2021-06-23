package main

import (
	"fmt"
)

func controlStructureTest() {

	f := true
	flag := &f

	// pointer을 선호한다함

	if flag == nil {
		fmt.Println("null")
	} else if !*flag {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	// for 문 type
	arr := []string{"my", "name", "is"}
	for i, value := range arr {
		fmt.Println(i, arr[i])
		fmt.Println(value)
	}

	// for in type
	myMap := make(map[string]interface{})
	myMap["name"] = "test"
	myMap["age"] = 20

	for k, v := range myMap {
		fmt.Printf("key: %s, value: %v\n", k, v)
	}

	// continue, break, switch

	day := "Fri"
	switch day {
	case "Fri":
		fmt.Println("Friday")
	case "Mon", "Tue", "Wed":
		fmt.Println("not Friday")
	default:
		fmt.Println("default")
	}

	switch {
	case day == "Fri":
		fmt.Println("TGIF")

	}

	// web에서는 다음과 같은 방식으로
	// err := errors.New("")
	// switch err {
	// case "error code":
	// 	fmt.Println("test")
	// }

}
