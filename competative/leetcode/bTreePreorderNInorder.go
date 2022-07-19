package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var orderMap map[int]int
var preorderIndex int

func main() {
	pre := []int{3, 9, 20, 15, 7}
	in := []int{9, 3, 15, 20, 7}
	tree := buildTree(pre, in)
	fmt.Println(tree)
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	orderMap = make(map[int]int)
	for i, v := range inorder {
		orderMap[v] = i
	}
	preorderIndex = 0
	return builder(preorder, 0, len(preorder)-1)
}

func builder(preorder []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	// select the preorder_index element as the root and increment it
	rootValue := preorder[preorderIndex]
	preorderIndex++
	root := &TreeNode{Val: rootValue}

	// build left and right subtree
	// excluding inorderIndexMap[rootValue] element because it's the root
	root.Left = builder(preorder, left, orderMap[rootValue]-1)
	root.Right = builder(preorder, orderMap[rootValue]+1, right)
	return root
}
