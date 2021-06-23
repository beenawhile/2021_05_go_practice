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

	root := &Node{val: 0}

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

	PrintNodes(root)

	// 2. remove
	// - 지울 때 주의점
	//  - 양 끝은?
	// 맨 앞
	//  - root 만 바꿔주면 됨
	// 맨 끝
	//  - tail 만 바꿔주면 됨
	root, tail = RemoveNode(tail, root, tail)

	PrintNodes(root)

	root, tail = RemoveNode(tail, root, tail)

	PrintNodes(root)

	root, tail = RemoveNode(root, root, tail)

	PrintNodes(root)
}

func PrintNodes(root *Node) {
	node := root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}

	fmt.Printf("%d\n", node.val)
}

// 1. Add

// func AddNode(root *Node, val int) {
// 	var tail *Node
// 	tail = root
// 	for tail.next != nil {
// 		tail = tail.next
// 	}

// 	node := &Node{val: val}
// 	tail.next = node
// }

// 해당 방법은 tail에 대한 정보가 없기 때문에 맨끝에 값을 추가하기 위해 for 문을 돌아야함.
//  => O(N)번 for문을 거치게 된다

func AddNode(tail *Node, val int) *Node {

	node := &Node{val: val}
	tail.next = node
	return node

}

// 해당방법은 node 하나만 추가해서 tail을 바꿔주기만 하면 됨. for 문을 안돌아도 됨
// => O(1)번 for문을 거치게 된다
// => 시간이 적게 걸림

// 2. Remove
// - 어떻게 할까? => 삭제 할 노드 전의 노드와 삭제할 노드 후의 노드를 이어주기만 하면 됨
// - 참조하고 있지 않은 노드는 reference count = 0 이기 때문에 메모리에서 사라짐
func RemoveNode(node, root, tail *Node) (*Node, *Node) {
	// 지울 때 주의점
	//  - 양 끝은?
	// 1. 맨 앞
	//  - root 만 바꿔주면 됨
	// 2. 맨 끝
	//  - tail 만 바꿔주면 됨
	// prev.next = prev.next.next
	// 3. node 가 하나일 때?
	// - root, tail 모두 nil 가리켜야함

	// 지우고자 하는 node가 첫번째 일 때
	if node == root {
		root = root.next

		// node 가 하나 밖에 없을 때
		if root == nil {
			tail = nil
		}

		return root, tail
	}

	// 지우고자 하는 node 전 node 찾기
	prev := root
	for prev.next != node {
		prev = prev.next
	}

	// node 찾았는데 tail node 일 경우
	if node == tail {
		prev.next = nil
		tail = prev
	} else {
		// tail node가 아닐 경우 : 전 것과 다음거만 이어주면 됨
		prev.next = prev.next.next
	}

	return root, tail

}

// - 내가 지우고 싶은 node를 찾기 위해 O(N) 만큼 반복해야 함
// 빠르게 하기 위해서 Double Linked List를 사용하면 됨(=prev도 알고 next도 알면 됨) => O(1)이 됨

// 비고법 : 알고리즘에서 시간을 나타내는 방법
// - element의 갯수와 알고리즘의 속도가 바뀔 때, 서로간의 상관관계
