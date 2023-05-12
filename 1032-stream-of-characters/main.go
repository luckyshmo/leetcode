package main

import "fmt"

func main() {
	words := []string{"cd", "f", "kl"}
	letter := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l'}
	obj := Constructor(words)
	for _, l := range letter {
		fmt.Println(obj.Query(l))
	}
}

type StreamChecker struct {
	chars string
	root  *Node
}

func Constructor(words []string) StreamChecker {
	return StreamChecker{
		root: New(words),
	}
}

func (sc *StreamChecker) Query(letter byte) bool {
	sc.chars += string(letter)
	cn := sc.root
	for i := len(sc.chars) - 1; i >= 0; i-- {
		node, _ := contains(cn.leafs, sc.chars[i])
		if node == nil {
			return false
		}
		if node.isLeaf {
			return true
		}
		if len(node.leafs) == 0 {
			return true
		}
		cn = node
	}
	return false
}

type Node struct {
	letter byte
	leafs  []*Node
	isLeaf bool
}

func New(words []string) (root *Node) {
	root = &Node{leafs: make([]*Node, 0)}
	temp := root
	for _, w := range words {
		for i := len(w) - 1; i >= 0; i-- {
			leaf, ok := contains(temp.leafs, w[i])
			if !ok {
				leaf = &Node{w[i], make([]*Node, 0), false}
				temp.leafs = append(temp.leafs, leaf)
			}
			if i == 0 {
				leaf.isLeaf = true
			}
			temp = leaf
		}
		temp = root
	}

	return root
}

func contains(leafs []*Node, letter byte) (*Node, bool) {
	for _, leaf := range leafs {
		if leaf.letter == letter {
			return leaf, true
		}
	}
	return nil, false
}
