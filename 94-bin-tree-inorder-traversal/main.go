package main

import (
	"fmt"
)

func main() {
	fmt.Println()
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	l := inorderTraversal(root.Left)
	l = append(l, root.Val)
	r := inorderTraversal(root.Right)
	return append(l, r...)
}
