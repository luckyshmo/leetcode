package helpers

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

func (t *TreeNode) Print() {
	stack := []*TreeNode{t}
	var currNode *TreeNode
	for len(stack) > 0 {
		currNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		println(currNode.Val)
		if currNode.Right != nil {
			stack = append(stack, currNode.Right)
		}
		if currNode.Left != nil {
			stack = append(stack, currNode.Left)
		}
	}
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
