package bintree

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

type Tree[Key constraints.Ordered, Data any] struct {
	root *Node[Key, Data]
}

func (t *Tree[Key, Data]) GetRoot() *Node[Key, Data] {
	return t.root
}

func (t *Tree[Key, Data]) Insert(key Key, data Data) {
	t.root = t.root.insert(key, data)
	if t.root.balanceFactor() < -1 || t.root.balanceFactor() > 1 {
		t.rebalance()
	}
}

func (t *Tree[Key, Data]) rebalance() {
	if t == nil || t.root == nil {
		return
	}
	t.root = t.root.rebalance()
}

func (t *Tree[Key, Data]) Find(k Key) (Data, bool) {
	if t == nil || t.root == nil {
		return *new(Data), false
	}
	return t.root.find(k)
}

func (t *Tree[Key, Data]) Delete(k Key) *Node[Key, Data] {
	if t == nil || t.root == nil {
		return nil
	}
	t.root = t.root.delete(k).rebalance()
	return t.root
}

func (t *Tree[Key, Data]) Traverse(n *Node[Key, Data], f func(*Node[Key, Data])) {
	if n == nil {
		return
	}
	t.Traverse(n.left, f)
	f(n)
	t.Traverse(n.right, f)
}

func (t *Tree[Key, Data]) PrettyPrint() {
	printNode := func(n *Node[Key, Data], depth int) {
		fmt.Printf("%s%v", strings.Repeat("    ", depth), n.key)
		fmt.Print("(", n.height, ")\n")
	}
	var walk func(*Node[Key, Data], int)
	walk = func(n *Node[Key, Data], depth int) {
		if n == nil {
			return
		}
		walk(n.right, depth+1)
		printNode(n, depth)
		walk(n.left, depth+1)
	}

	walk(t.root, 0)
}

func (t *Tree[Value, Data]) Dump() {
	t.root.dump(0, "")
}
