package dataStruct

import "fmt"

// Double Linked List
// - linked list 에서 remove 시 발생하였던 O(N) -> O(1)로 줄일 수 있음
// - Next 뿐만아니라 Prev 정보까지 저장

type Node struct {
	Prev *Node
	Next *Node
	Val  int
}

type DoubleLinkedList struct {
	Root *Node
	Tail *Node
}

func (l *DoubleLinkedList) AddNode(Val int) {
	if l.Root == nil {
		l.Root = &Node{Val: Val}
		l.Tail = l.Root
		return
	}

	l.Tail.Next = &Node{Val: Val}
	Prev := l.Tail
	l.Tail = l.Tail.Next
	l.Tail.Prev = Prev
}

func (l *DoubleLinkedList) RemoveNode(node *Node) {
	if node == l.Root {
		l.Root = node.Next
		if l.Root != nil {
			l.Root.Prev = nil
		}
		node.Next = nil
		return
	}

	Prev := node.Prev

	if node == l.Tail {
		Prev.Next = nil
		l.Tail.Prev = nil
		l.Tail = Prev
	} else {
		node.Prev = nil
		Prev.Next = Prev.Next.Next
		Prev.Next.Prev = Prev
	}
	node.Next = nil
}

func (l *DoubleLinkedList) Back() int {
	if l.Tail != nil {
		return l.Tail.Val
	}
	return 0
}

func (l *DoubleLinkedList) Front() int {
	if l.Root != nil {
		return l.Root.Val
	}
	return 0
}

func (l *DoubleLinkedList) PopBack() {
	if l.Tail == nil {
		return
	}
	l.RemoveNode(l.Tail)
}

func (l *DoubleLinkedList) PopFront() {
	if l.Root == nil {
		return
	}
	l.RemoveNode(l.Root)
}

func (l *DoubleLinkedList) Empty() bool {
	return l.Root == nil
}

func (l *DoubleLinkedList) PrintNodes() {
	node := l.Root
	for node.Next != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}

	fmt.Printf("%d\n", node.Val)
}

func (l *DoubleLinkedList) PrintReverse() {
	node := l.Tail
	for node.Prev != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Prev
	}
	fmt.Printf("%d\n", node.Val)
}
