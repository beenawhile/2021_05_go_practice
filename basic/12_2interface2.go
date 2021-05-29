// interface : 두 객체간에 어떤 관계를 가지는지 정의한 것
// - object에서 기능을 외부로 빼냈다고 생각해도 된다함

package main

import "fmt"

// interface 기능 : Generic 같은것인듯?
// - 다시보니 dynamic 에 가까운 듯
func Anything(anything interface{}) {
	fmt.Println(anything)
}

func interfaceTest2() {
	Anything(2.34)
	Anything("Test")
	Anything(2)
	Anything(struct{}{})

	// map 형태를 만드는데 key는 string, value는 anything
	myMap := make(map[string]interface{})
	myMap["name"] = "test"
	myMap["age"] = 10
	fmt.Println(myMap)

}
