package main

import (
	"dataStruct"
	"fmt"
)

// binary tree : 자식이 두 개 까지만 가능한 tree
//  - 자식을 갯수가 정해지지 않은 slice 형태가 아닌 left, right 2개를 가지고 있음
//  - 검색을 빨리 하기 위해서 사용함

func binaryTreeTest() {
	tree := dataStruct.NewBinaryTree(5)

	tree.Root.AddNode(3)
	tree.Root.AddNode(2)
	tree.Root.AddNode(4)
	tree.Root.AddNode(8)
	tree.Root.AddNode(7)
	tree.Root.AddNode(6)
	tree.Root.AddNode(10)
	tree.Root.AddNode(9)

	tree.Print()

	fmt.Println()

	// 초기문 :if 함수를 먼저 수행하고 결과를 이용해서 사용
	if found, cnt := tree.Search(6); found {
		fmt.Println("found 6. cnt: ", cnt)
	} else {
		fmt.Println("not found 6. cnt: ", cnt)
	}

	if found, cnt := tree.Search(12); found {
		fmt.Println("found 12. cnt: ", cnt)
	} else {
		fmt.Println("not found 12. cnt: ", cnt)
	}

}
