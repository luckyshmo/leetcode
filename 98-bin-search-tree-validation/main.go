package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// n4 := &TreeNode{
	// 	Val: 4,
	// }
	// n6 := &TreeNode{
	// 	Val: 6,
	// }
	// n0 := &TreeNode{
	// 	Val: 0,
	// }
	// n2 := &TreeNode{
	// 	Val: 2,
	// }
	// n1 := &TreeNode{
	// 	Val:   1,
	// 	Left:  n0,
	// 	Right: n2,
	// }
	// n5 := &TreeNode{
	// 	Val:   5,
	// 	Left:  n4,
	// 	Right: n6,
	// }
	// n3 := &TreeNode{
	// 	Val:   3,
	// 	Left:  n1,
	// 	Right: n5,
	// }

	// fmt.Println(isValidBST(n3))

	//! [5,4,6,null,null,3,7] false
	n3 := &TreeNode{
		Val: 3,
	}
	n7 := &TreeNode{
		Val: 7,
	}
	n6 := &TreeNode{
		Val:   6,
		Left:  n3,
		Right: n7,
	}
	n4 := &TreeNode{
		Val: 4,
	}
	n5 := &TreeNode{
		Val:   5,
		Left:  n4,
		Right: n6,
	}
	println(isValidBST(n5))

	//! [2,1,3] true
	// n1 := &TreeNode{
	// 	Val: 1,
	// }
	// n3 := &TreeNode{
	// 	Val: 3,
	// }
	// n2 := &TreeNode{
	// 	Val:   2,
	// 	Left:  n1,
	// 	Right: n3,
	// }
	// println(isValidBST(n2))
}

func isValidBST(root *TreeNode) bool {
	return rec(root, math.MinInt64, math.MaxInt64)
}

func rec(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.Val >= max || root.Val <= min {
		return false
	}
	return rec(root.Left, min, root.Val) && rec(root.Right, root.Val, max)
}
