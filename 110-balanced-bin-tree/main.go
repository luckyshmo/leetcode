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

func (t *TreeNode) PrettyPrint() {
	printNode := func(n *TreeNode, depth int) {
		fmt.Printf("%s%v", strings.Repeat("    ", depth), n.Val)
		fmt.Print("\n")
	}
	var walk func(*TreeNode, int)
	walk = func(n *TreeNode, depth int) {
		if n == nil {
			return
		}
		walk(n.Right, depth+1)
		printNode(n, depth)
		walk(n.Left, depth+1)
	}

	walk(t, 0)
}

func main() {
	t2 := NewTreeNode("1,2,3,4,5,6,null,8")
	t1 := NewTreeNode("1,2,2,3,3,null,null,4,4")
	t1.PrettyPrint()
	fmt.Println(isBalanced(t1))
	fmt.Println(isBalanced(t2))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBalancedR(root) > 0
}

func isBalancedR(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := isBalancedR(root.Left)
	if lh == -1 {
		return -1
	}
	rh := isBalancedR(root.Right)
	if rh == -1 {
		return -1
	}

	if abs(lh-rh) > 1 {
		return -1
	}

	if lh > rh {
		return lh + 1
	}
	return rh + 1
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
