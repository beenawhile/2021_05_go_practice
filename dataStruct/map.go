// map = dictionary
//  - key-value 타입으로 이루어진 자료구조
//  - 단순히 늘어놓기만 하면 read 시에 O(N) 의 속도
//  - BST(Binary Search Tree)를 이용하여 속도를 높이는 방법을 사용함 => sorted map, ordered map
//  	- 속도 : log2N

// hash map
//  - 같은 입력 -> 같은 출력
//  - 다른 입력 -> 다른 출력
//  - 배열에 값을 만들고, hash 라는 함수에 key 값을 넣어 나온 인덱스 값을 이용해 배열에서 해당 인덱스 값을 불러옴
//  - 속도 : O(1) => 일정한 속도 => 시간이 거의 걸리지 않음

// * 관건
//  - hash 를 어떻게 만드는가?
//    1. 출력값의 범위가 정해져 있음
//    2. 같은 입력이면 같은 출력이 나옴
//    3. 다른 입력이면 (보통의 경우) 다른 출력이 나옴
//   - sine : 출력값 범위 정해져 있음
//   - mod : 출력값 범위 정해져 있음, 연산 단순함 => hash에 많이 사용된다 함
//     - one-way function : 나눔값과 나머지를 통해 원래 값을 추론할 수 없음(추론값이 너무 많기 때문) => 한 방향으로만 계산할 수 있기 때문에 one way function
//     - 즉, 암호문을 만들 수 있다는 것임

// Rolling hash
// - 한글자씩 구르면서 만들어진다고 해서 rolling hash
// - "abcde"
//   s0s1s2s3s4 라고 놓았을 때
// - S0...Sn,  Hi = (Hi-1 * A + Si) % B
// - A : S값의 범위(문자의 경우 ASCII가 0~255 이기 때문에 256)
// - B : 소수(값의 분포가 넓어지기 때문에). 여기서는 3571

package dataStruct

func Hash(s string) int {
	h := 0
	A := 256
	B := 3571
	for i := 0; i < len(s); i++ {
		h = (h*A + int(s[i])) % B
	}
	return h

}

// 문제
//  - 다른 key 인데 같은 hash 값을 가져 값이 충돌나는 경우가 생길 수 있음
//  - 단순 방지 방법
//    - array 안에 또 array를 집어 넣어, array를 돌면서 같은 값을 찾음(hash가 같은 것만 돌기 때문에 전체를 도는 것보다는 매우 효율적)

type KeyValue struct {
	Key   string
	Value string
}

type Map struct {
	keyArray [3571][]KeyValue
}

func CreateMap() *Map {
	return &Map{}
}

func (m *Map) Add(key, value string) {
	h := Hash(key)
	m.keyArray[h] = append(m.keyArray[h], KeyValue{Key: key, Value: value})
}

func (m *Map) Get(key string) string {
	h := Hash(key)
	for i := 0; i < len(m.keyArray[h]); i++ {
		if m.keyArray[h][i].Key == key {
			return m.keyArray[h][i].Value
		}
	}
	return ""
}

// 장점 : 빠르다
// 단점 : 해쉬의 결과값이 정렬되어 있는 상태가 아니기 때문에 정렬된 상태로 뽑아낼 수 없음
