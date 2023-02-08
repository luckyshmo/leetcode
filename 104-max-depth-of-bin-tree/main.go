package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println()
}

func maxDepth(root *TreeNode) (maxH int) {
	if root == nil {
		return
	}

	type Tuple struct {
		n *TreeNode
		h int
	}

	queue := []Tuple{{root, 0}}
	var currNode Tuple
	for len(queue) > 0 {
		currNode = queue[0]
		queue = queue[1:]

		if currNode.h > maxH {
			maxH = currNode.h
		}

		if currNode.n.Right != nil {
			queue = append(queue, Tuple{currNode.n.Right, currNode.h + 1})
		}
		if currNode.n.Left != nil {
			queue = append(queue, Tuple{currNode.n.Left, currNode.h + 1})
		}
	}
	return maxH + 1
}
