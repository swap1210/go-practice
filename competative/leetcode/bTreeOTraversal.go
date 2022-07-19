package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	fmt.Println(levelOrder(root))
	// x := []int{12, 3, 4}
	// fmt.Println(x)
	// x = x[1:]
	// fmt.Println(x)
}

func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		row := []int{}
		for i := 0; i < l; i++ {
			cur := queue[0]
			// fmt.Println("b", queue)
			queue = queue[1:]
			// fmt.Println("a", queue)
			row = append(row, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		ans = append(ans, row)
	}

	return ans
}
