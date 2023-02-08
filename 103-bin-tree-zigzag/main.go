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
	t := NewTreeNode("1,2,3,4,null,null,5")
	t.PrettyPrint()
	fmt.Println(zigzagLevelOrder(t))
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

func zigzagLevelOrder(root *TreeNode) (layers [][]int) {
	if root == nil {
		return nil
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

		if len(layers) <= currNode.h {
			layers = append(layers, []int{currNode.n.Val})
		} else {
			if currNode.h%2 == 0 {
				layers[currNode.h] = append([]int{currNode.n.Val}, layers[currNode.h]...)
			} else {
				layers[currNode.h] = append(layers[currNode.h], currNode.n.Val)
			}
		}

		if currNode.n.Right != nil {
			queue = append(queue, Tuple{currNode.n.Right, currNode.h + 1})
		}
		if currNode.n.Left != nil {
			queue = append(queue, Tuple{currNode.n.Left, currNode.h + 1})
		}

	}
	return
}
