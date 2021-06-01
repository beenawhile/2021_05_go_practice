// linked list : 연결된 리스트?

//  - array 구조
// ==========================================
//      |data|data|data|data|data|
// ==========================================
//      - 하나의 덩어리 메모리를 할당해서 나눠서 사용함 (붙어있음)

//  - linked list 구조
// ==========================================
//     -|data|-|data|-|data|-|data|-|data|-
// ==========================================
//      - 붙어있지 않고 연결해놓음 => 메모리 상에서 떨어져 있을 수 있음
//      - 각각의 노드를 어떻게 연결하나? => pointer로 연결
//      - structure를 정의하고 다음 노드에 대한 포인터를 structure의 한 요소로 가지고 있음
// ==========================================
// type Node struct {
// 	next *Node
// 	val int
// }
// 시작하는 노드(=root node)에 대한 정보는 알고 있어야 함

package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func linkedListTest() {
	var root *Node

	root = &Node{val: 0}

	// element 추가는 어떻게 해야하나?
	// 하나의 노드 만든 다음에 맨끝의 Node next 값을 새로 만든 노드로 갈아 끼우면 됨
	// 맨끝의 노드는 어떤건지 어떻게 알 수 있나?

	//  1. root 부터 시작해서 다음 next가 없을 때 까지 찾아감

	// var tail *Node
	// tail = root
	// for tail.next != nil {
	// 	tail = tail.next
	// }

	// node := &Node{val: 1}
	// tail.next = node
	// 이 추가하는 과정을 function으로 만듬

	// for i := 1; i < 10; i++ {
	// 	AddNode(root, i)
	// }

	// // 출력
	// node := root
	// for node.next != nil {
	// 	fmt.Printf("%d -> ", node.val)
	// 	node = node.next
	// }

	// fmt.Printf("%d\n", node.val)

	//  2. tail을 항상 가지고 기억하고 있으면 됨 => tail을 새로 추가된 노드로 바꿔주기만 하면 됨
	var tail *Node
	// 처음에는 root 이면서 tail 이 됨
	tail = root
	for i := 1; i < 10; i++ {
		tail = AddNode(tail, i)
	}

	node := root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}

	fmt.Printf("%d\n", node.val)
}

// func AddNode(root *Node, val int) {
// 	var tail *Node
// 	tail = root
// 	for tail.next != nil {
// 		tail = tail.next
// 	}

// 	node := &Node{val: val}
// 	tail.next = node
// }

func AddNode(tail *Node, val int) *Node {

	node := &Node{val: val}
	tail.next = node
	return node

}
