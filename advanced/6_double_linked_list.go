package main

import "fmt"

// Double Linked List
// - linked list 에서 remove 시 발생하였던 O(N) -> O(1)로 줄일 수 있음
// - next 뿐만아니라 prev 정보까지 저장

type Point struct {
	prev *Point
	next *Point
	val  int
}

type DoubleLinkedList struct {
	root *Point
	tail *Point
}

func (l *DoubleLinkedList) AddNode(val int) {
	if l.root == nil {
		l.root = &Point{val: val}
		l.tail = l.root
		return
	}

	l.tail.next = &Point{val: val}
	prev := l.tail
	l.tail = l.tail.next
	l.tail.prev = prev
}

func (l *DoubleLinkedList) RemoveNode(node *Point) {
	if node == l.root {
		l.root = node.next
		l.root.prev = nil
		node.next = nil
		return
	}

	prev := node.prev

	if node == l.tail {
		prev.next = nil
		l.tail.prev = nil
		l.tail = prev
	} else {
		node.prev = nil
		prev.next = prev.next.next
		prev.next.prev = prev
	}
	node.next = nil
}

func (l *DoubleLinkedList) PrintNodes() {
	node := l.root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}

	fmt.Printf("%d\n", node.val)
}

func (l *DoubleLinkedList) PrintReverse() {
	node := l.tail
	for node.prev != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.prev
	}
	fmt.Printf("%d\n", node.val)
}

func doubleLinkedListTest() {
	list := &DoubleLinkedList{}
	list.AddNode(0)

	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}

	list.PrintNodes()

	list.RemoveNode(list.root.next)

	list.PrintNodes()

	list.RemoveNode(list.root)

	list.PrintNodes()

	list.RemoveNode(list.tail)

	list.PrintNodes()

	fmt.Printf("tail:%d\n", list.tail.val)

	fmt.Println("거꾸로 출력")

	list.PrintReverse()

	// slice 와 Double Linked List 비교

	// 1. Add
	// slice : capacity를 넘는 append 시도할 때 새 슬라이스 공간을 확보한 다음 값을 복사해서 이동시킴 => O(N)
	// Linked List : 하나만 추가해주면 됨 O(1)
	// Double linked list : 하나만 추가해주면 됨 O(1)

	// 2. Remove
	// slice :
	//  - 끝의 것을 삭제 :하나만 reference count를 감소시켜 주면 됨 => O(1)
	//  - 중간 것을 삭제 :새 공간을 확보한 다음 삭제할 것만 빼고 복사 =>
	a := []int{1, 2, 3, 4, 5}
	a = append(a[:2], a[3:]...)
	fmt.Println(a, len(a), cap(a))
	//    cap 결과를 확인해보면 뒤에 것을 자르고 다시 for문을 돌면서 붙여넣는 것을 알 수 있음 => O(N)
	//  - 맨 앞을 삭제 :시작 인덱스만 바꿔주면 됨 => O(1)
	// Linked list : prev로 찾아가서 이어줘야함 O(N)
	// Double linked list : prev 이어주기만 하면 되기 때문에 O(1)

	// ※ 무조건 list가 좋나? => 아님

	// 3. Read (Random Access)
	// slice : 연속된 메모리 이기 때문에 특정 값을 가져올 때 바로 가져올 수 있음(시작 메모리 주소, 각 항목 메모리 사이즈 알고 있기 때문) => O(1)
	// list : 특정 노드 까지 찾아가아함=> O(N)

	// 4. 하드웨어적 비교
	// slice : cache는 메모리 덩어리(근방에 있는 것들)를 들고와서 연산하기 때문에 slice 특성상(덩어리 형태로 되어 있음) 메모리에 한번 더 안가도 됨
	//  -> cache 활용에 장점이 있음
	// list : 메모리 상에 떨어진 장소에 적재되어있기 때문에 cache가 여러번 가져오게 됨
	//  -> 수량이 엄청 많아지게 되면 cache miss(=memory에서 cache에 가져온 것이 아무 쓸모 없음)가 계속 발생하여 CPU가 부담이 되고 느려짐

	// 총평
	// - 적은 element가 있을 때는 크게 차이가 없음
	// - 데이터 크기가 커져서 많은 수가 있다 => linked list 고려해볼만 함
	// - 데이터 집약도(memory에 몰려있는 정도)가 중요하다 => slice 고려

}
