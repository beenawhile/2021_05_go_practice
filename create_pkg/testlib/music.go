package testlib

import "fmt"

// 사이즈가 큰 복잡한 라이브러리 같은 경우, "go install" 명령을 사용하여 라이브러리를 컴파일하여 Cache할 수 있음

var pop map[string]string

func init() {
	pop = make(map[string]string)
	pop["Adele"] = "Hello"
	pop["Alicia Keys"] = "Fallin'"
	pop["John Legend"] = "All of Me"
}

// GetMusic : Popular music by singer (외부에서 호출 가능)
func GetMusic(singer string) string {
	return pop[singer]
}

func getKeys() { // 내부에서만 호출 가능
	for _, kv := range pop {
		fmt.Println(kv)
	}
}
