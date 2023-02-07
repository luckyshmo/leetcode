package main

import "fmt"

func main() {
	//! [1,2,2,null,3,null,3] false
	n3 := &TreeNode{
		Val: 3,
	}
	n33 := &TreeNode{
		Val: 3,
	}
	n2 := &TreeNode{
		Val:   2,
		Right: n3,
	}
	n22 := &TreeNode{
		Val:   2,
		Right: n33,
	}
	n1 := &TreeNode{
		Val:   1,
		Left:  n2,
		Right: n22,
	}
	println(isSymmetric(n1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left != nil && root.Right == nil {
		return false
	}
	if root.Left == nil && root.Right != nil {
		return false
	}
	return inorderTraversalRL(root.Right) == inorderTraversalLR(root.Left)
}

func inorderTraversalLR(root *TreeNode) (res string) {
	stack := []*TreeNode{root}
	var currNode *TreeNode
	for len(stack) > 0 {
		currNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res += fmt.Sprint(currNode.Val)
		if currNode.Right != nil {
			res += "r"
			stack = append(stack, currNode.Right)
		}
		if currNode.Left != nil {
			res += "l"
			stack = append(stack, currNode.Left)
		}
	}
	return
}

func inorderTraversalRL(root *TreeNode) (res string) {
	stack := []*TreeNode{root}
	var currNode *TreeNode
	for len(stack) > 0 {
		currNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res += fmt.Sprint(currNode.Val)
		if currNode.Left != nil {
			res += "r"
			stack = append(stack, currNode.Left)
		}
		if currNode.Right != nil {
			res += "l"
			stack = append(stack, currNode.Right)
		}
	}
	return
}
