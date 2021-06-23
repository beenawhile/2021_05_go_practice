package dataStruct

import "fmt"

type TreeNode struct {
	Val      int
	Children []*TreeNode
}

type Tree struct {
	Root *TreeNode
}

func (t *Tree) AddNode(val int) {
	if t.Root == nil {
		t.Root = &TreeNode{Val: val}
	} else {
		t.Root.Children = append(t.Root.Children, &TreeNode{Val: val})
	}
}

func (t *TreeNode) AddNode(val int) {
	t.Children = append(t.Children, &TreeNode{Val: val})
}

// 데이터를 불러오는 방법은
// 1. DFS(Depth-First Search) : 깊이 우선 탐색
// - 만드는 방법
//  1) 재귀호출 recursive call
//  2) 스택 stack
// 2. BFS(Breadth-First Search) : 깊이 우선 탐색
//  1) 재귀호출 recursive call >> 너무 어려워 잘 안씀
//  2) 큐 Queue

// 1. DFS
func (t *Tree) DFS1() {
	// 1. 재귀호출 방법
	DFS1(t.Root)
}

func (t *Tree) DFS2() {
	// 2. Stack 을 이용한 방법
	// 만들어놓은 Linked List를 전반적으로 수정해야하기 때문에 slice로 만듬
	s := []*TreeNode{}
	s = append(s, t.Root)

	for len(s) > 0 {
		var last *TreeNode
		last, s = s[len(s)-1], s[:len(s)-1]

		fmt.Printf("%d->", last.Val)

		for i := len(last.Children) - 1; i >= 0; i-- {
			s = append(s, last.Children[i])
		}
	}
}

func DFS1(node *TreeNode) {
	fmt.Printf("%d->", node.Val)

	for i := 0; i < len(node.Children); i++ {
		DFS1(node.Children[i])
	}
}

// BFS
func (t *Tree) BFS() {
	queue := []*TreeNode{}
	queue = append(queue, t.Root)
	for len(queue) > 0 {
		var first *TreeNode
		first, queue = queue[0], queue[1:]
		fmt.Printf("%d->", first.Val)

		for i := 0; i < len(first.Children); i++ {
			queue = append(queue, first.Children[i])
		}
	}
}

// DFS 와 BFS를 어디다 쓸 수 있는가?
// - Tree는 많은 곳에 사용하고 있음 => directory 구조
// - 게임에서의 예제 : 길찾기
//   - Dijkstra Algorithm : 최단거리 찾는 알고리즘
//     - 정확하게 tree는 아니고 graph (graph : 노드와 노드 사이를 연결한 구조, tree, linked list도 graph의 한 종류)
//     - 시작점에서 끝점까지 가는 최단거리 찾는 알고리즘 => DFS를 이용하는 알고리즘
//   - 게임에서는 이 알고리즘을 수정한 A* 알고리즘을 사용함
