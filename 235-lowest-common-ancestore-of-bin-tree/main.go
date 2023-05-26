package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// n1 := TreeNode{
	// 	Val: 1,
	// }
	// n3 := TreeNode{
	// 	Val: 3,
	// }
	// n2 := TreeNode{
	// 	Val:   2,
	// 	Left:  &n1,
	// 	Right: &n3,
	// }

	// fmt.Println(lowestCommonAncestor(&n2, &n1, &n3))

	tr := NewTreeNode("6,2,8,0,4,7,9,null,null,3,5")
	f1 := tr.Left.Right.Left
	f2 := tr.Left.Right.Right
	fmt.Println(lowestCommonAncestor(tr, f1, f2))

	// tr := NewTreeNode("5,3,6,2,4,null,null,1")
	// f1, f2 := tr.Left.Left.Left, tr.Left.Right
	// fmt.Println(lowestCommonAncestor(tr, f1, f2))
}

var found = false

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	found = false
	return helper(root, p, q)
}

func helper(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	this := root == p || root == q

	left := helper(root.Left, p, q)
	if found {
		return left
	}
	fInLeft := left == p || left == q
	if fInLeft && this {
		found = true
		return root
	}

	right := helper(root.Right, p, q)
	if found {
		return right
	}
	fInRight := right == p || right == q
	if fInRight && this {
		found = true
		return root
	}
	if found {
		return right
	}

	if fInLeft && fInRight {
		found = true
		return root
	}

	if left == nil && right == nil && !this {
		return nil
	} else if left != nil {
		return left
	} else if right != nil {
		return right
	}

	return root
}

func NewTreeNode(str string) (root *TreeNode) {
	elems := strings.Split(str, ",")

	layers := make([][]*TreeNode, 0)
	layerLength := 1
	for i := 0; i < len(elems); {
		layer := make([]*TreeNode, layerLength)
		for j := 0; j < layerLength && i < len(elems); j, i = j+1, i+1 {
			if elems[i] == "null" {
				layer[j] = nil
			} else {
				el, err := strconv.Atoi(elems[i])
				if err != nil {
					panic(err)
				}
				layer[j] = &TreeNode{
					Val: el,
				}
				if len(layers) > 0 {
					if j%2 == 0 {
						layers[len(layers)-1][j/2].Left = layer[j]
					} else {
						layers[len(layers)-1][j/2].Right = layer[j]
					}
				}
			}
		}
		layers = append(layers, layer)
		layerLength = layerLength * 2
	}
	return layers[0][0]
}
