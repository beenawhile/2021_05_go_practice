package dataStruct

import "fmt"

// binary tree : 자식이 두 개 까지만 가능한 tree
//  - 자식을 갯수가 정해지지 않은 slice 형태가 아닌 left, right 2개를 가지고 있음

type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

// 왜 중요한가? => BST(Binary Search Tree) 때문에
// - 모든 노드에 left는 부모 노드보다 작은 값, right는 부모 노드 보다 큰 값이 오는 트리
// - tree에서 어떤 값을 찾을 때, 다른 search (BFS, DFS)는 모든 노드를 거쳐가야 하지만 BST는 모든 노드를 다 검사할 필요가 없음
// - 노드 값을 비교해보고 찾고자 하는 값이 큰값인지 작은값인지에 따라 한 부분만 찾아보면 됨
// - big O 법으로 속도를 계산할 때 속도 : log2N (다른 tree 구조는 O(N))
// - BST는 대칭을 이뤄야 검색 효율이 더 좋아짐 (머릿속으로 잠시 생각해보면 이해할 수 있을 듯)
//  => 높이(=depth)가 가장 낮은 tree 가 가장 효율적 => 최소 신장 트리
// - 이런 tree를 어떻게 만드나? => 기존 트리를 회전시켜 최소 신장을 만드는 방법 사용 => AVL 트리 알고리즘
//   다른 방법도 있음. Black Red tree, etc ...

type BinaryTree struct {
	Root *BinaryTreeNode
}

func NewBinaryTree(v int) *BinaryTree {
	tree := &BinaryTree{}
	tree.Root = &BinaryTreeNode{Val: v}
	return tree
}

func (n *BinaryTreeNode) AddNode(v int) *BinaryTreeNode {
	if n.Val > v {
		if n.Left == nil {
			n.Left = &BinaryTreeNode{Val: v}
			return n.Left
		} else {
			return n.Left.AddNode(v)
		}
	} else {
		if n.Right == nil {
			n.Right = &BinaryTreeNode{Val: v}
			return n.Right
		} else {
			return n.Right.AddNode(v)
		}
	}
}

// 몇번째 층에 있는지 알기 위해 타입 정의
type depthNode struct {
	depth int
	node  *BinaryTreeNode
}

func (t *BinaryTree) Print() {
	// BFS를 이용하여 프린트
	q := []depthNode{}
	q = append(q, depthNode{depth: 0, node: t.Root})
	currentDepth := 0

	for len(q) > 0 {
		var first depthNode
		first, q = q[0], q[1:]

		if first.depth != currentDepth {
			fmt.Println()
			currentDepth = first.depth
		}
		fmt.Print(first.node.Val, " ")

		if first.node.Left != nil {
			q = append(q, depthNode{depth: currentDepth + 1, node: first.node.Left})
		}
		if first.node.Right != nil {
			q = append(q, depthNode{depth: currentDepth + 1, node: first.node.Right})
		}
	}
}

func (t *BinaryTree) Search(v int) (bool, int) {
	return t.Root.Search(v, 1)
}

func (n *BinaryTreeNode) Search(v int, cnt int) (bool, int) {
	if n.Val == v {
		return true, cnt
	} else if n.Val > v {
		if n.Left != nil {
			return n.Left.Search(v, cnt+1)
		}
		return false, cnt
	} else {
		if n.Right != nil {
			return n.Right.Search(v, cnt+1)
		}
		return false, cnt
	}
}
