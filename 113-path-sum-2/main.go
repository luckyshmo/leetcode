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

func buildTree(s string) *TreeNode {
	s = s[1 : len(s)-1]
	split := strings.Split(s, ",")
	if len(split) == 0 {
		return nil
	}
	var root *TreeNode
	if split[0] != "null" {
		rootVal, _ := strconv.Atoi(split[0])
		root = &TreeNode{Val: rootVal}
	}
	nodes := []*TreeNode{root}
	for i, c := range split[1:] {
		if c != "null" {
			val, _ := strconv.Atoi(c)
			node := &TreeNode{Val: val}
			if i&1 == 1 {
				nodes[i/2].Right = node
			} else {
				nodes[i/2].Left = node
			}
			nodes = append(nodes, node)
		} else {
			nodes = append(nodes, nil)
		}
	}
	return root
}

func buildTree1(s string) *TreeNode {
	s = s[1 : len(s)-1]
	split := strings.Split(s, ",")
	if len(split) == 0 {
		return nil
	}
	rootVal, _ := strconv.Atoi(split[0])
	root := &TreeNode{Val: rootVal}
	nodes := []*TreeNode{root}
	for i, c := range split[1:] {
		if c != "null" {
			val, _ := strconv.Atoi(c)
			node := &TreeNode{Val: val}
			if i&1 == 1 {
				nodes[i/2].Right = node
			} else {
				nodes[i/2].Left = node
			}
			nodes = append(nodes, node)
		} else {
			nodes = append(nodes, nil)
		}
	}
	return root
}

func main() {
	r := buildTree("[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]")
	r1 := buildTree("-260,-202,-903,-980,-570,-858,218,764,-300,205,null,-35,null,null,-204,950,-769,258,-652,614,-584,76,817,-192,null,null,-114,880,null,-200,71,671,344,801,193,-18,876,-920,-730,222,679,null,-680,null,null,null,-859,744,-261,692,null,-341,-163,null,null,482,-979,205,null,146,165,801,100,-656,714,-629,995,474,307,-581,-150,-941,null,null,null,-937,-69,-23,82,null,-139,-591,null,-453,-861,-370,null,null,null,216,233,null,430,null,5,-110,null,null,-660,624,-510,-588,null,null,381,null,368,559,null,521,-301,null,522,379,null,null,null,null,456,519,null,null,482,349,null,null,19,null,null,288,-811,null,-372,null,null,-536,null,-404,-457,-740,860,null,null,-636,null,null,342,-874,-462,-504,781,855,-392,null,null,null,406,null,-758,541,null,-947,null,null,null,null,null,-964,null,600,-45,null,null,null,null,null,null,null,null,null,-194,null,null,null,-802,null,null,null,-3,null,-792,672,643,null,14,null,null,489,457,null,null,null,null,412,null,558,null,null,null,null,-846,158,-146,null,null,-76,-650,null,-782,null,-127,null,-678,null,null,null,null,null,null,-464,-426,null,-366,null,null,null,null,null,81,-607,716,null,null,-213,null,379,null,null,null,null,644,445,null,null,-419,-845,-720,null,null,-915,null,null,null,null,null,null,-686,594,-243,null,496,null,907,null,null,null,null,null,null,579,873,702,null,null,null,-834,null,null,null,null,null,-300,-214,-466,null,null,972,null,null,null,814,null,-940,null,763,null,null,null,null,-449,-844,null,null,null,null,-47")
	r1.PrettyPrint()
	r.PrettyPrint()
	// a := pathSum(t, 22)
	// fmt.Printf("%v", a)
}

func pathSum(n *TreeNode, ts int) [][]int {
	if n == nil {
		return nil
	}
	return rec(n, ts, make([]int, 0))
}

func rec(n *TreeNode, ts int, ar []int) [][]int {
	if n.Left == nil && n.Right == nil {
		if ts == n.Val {
			return [][]int{append(ar, n.Val)}
		}
		return nil
	}
	var left, right [][]int
	if n.Left != nil {
		left = rec(n.Left, ts-n.Val, append(append([]int{}, ar...), n.Val))
	}
	if n.Right != nil {
		right = rec(n.Right, ts-n.Val, append(append([]int{}, ar...), n.Val))
	}

	return append(left, right...)
}
