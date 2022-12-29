package bintree

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

type Node[Key constraints.Ordered, Data any] struct {
	key    Key
	data   Data
	left   *Node[Key, Data]
	right  *Node[Key, Data]
	height int
}

func (n *Node[Key, Data]) GetData() Data {
	return n.data
}
func (n *Node[Key, Data]) GetKey() Key {
	return n.key
}

func (n *Node[Key, Data]) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *Node[Key, Data]) balanceFactor() int {
	return n.right.getHeight() - n.left.getHeight()
}

func (n *Node[Key, Data]) insert(key Key, data Data) *Node[Key, Data] {
	if n == nil {
		return &Node[Key, Data]{
			key:    key,
			data:   data,
			height: 1,
		}
	}
	if n.key == key {
		n.data = data
		return n
	}
	if key < n.key {
		n.left = n.left.insert(key, data)
	} else {
		n.right = n.right.insert(key, data)
	}
	n.height = max(n.left.getHeight(), n.right.getHeight()) + 1

	return n.rebalance()
}

func (n *Node[Key, Data]) rotateLeft() *Node[Key, Data] {
	r := n.right
	n.right = r.left
	r.left = n
	n.height = max(n.left.getHeight(), n.right.getHeight()) + 1
	r.height = max(r.left.getHeight(), r.right.getHeight()) + 1
	return r
}

func (n *Node[Key, Data]) rotateRight() *Node[Key, Data] {
	l := n.left
	n.left = l.right
	l.right = n
	n.height = max(n.left.getHeight(), n.right.getHeight()) + 1
	l.height = max(l.left.getHeight(), l.right.getHeight()) + 1
	return l
}

func (n *Node[Key, Data]) rotateRightLeft() *Node[Key, Data] {
	n.right = n.right.rotateRight()
	n = n.rotateLeft()
	n.height = max(n.left.getHeight(), n.right.getHeight()) + 1
	return n
}

func (n *Node[Key, Data]) rotateLeftRight() *Node[Key, Data] {
	n.left = n.left.rotateLeft()
	n = n.rotateRight()
	n.height = max(n.left.getHeight(), n.right.getHeight()) + 1
	return n
}

func (n *Node[Value, Data]) rebalance() *Node[Value, Data] {
	switch {
	case n.balanceFactor() < -1 && n.left.balanceFactor() == -1:
		return n.rotateRight()
	case n.balanceFactor() > 1 && n.right.balanceFactor() == 1:
		return n.rotateLeft()
	case n.balanceFactor() < -1 && n.left.balanceFactor() == 1:
		return n.rotateLeftRight()
	case n.balanceFactor() > 1 && n.right.balanceFactor() == -1:
		return n.rotateRightLeft()
	}
	return n
}

func (n *Node[Key, Data]) find(s Key) (Data, bool) {
	switch {
	case s == n.key:
		return n.data, true
	case s < n.key:
		return n.left.find(s)
	default:
		return n.right.find(s)
	}
}

/* Given a non-empty binary search tree,
return the node with minimum key value
found in that tree. Note that the entire
tree does not need to be searched. */
func (n *Node[Key, Data]) minValueNode() *Node[Key, Data] {
	current := n

	/* loop down to find the leftmost leaf */
	for n.left != nil {
		current = n.left
	}

	return current
}

func Max[T constraints.Ordered](arg1, arg2 T) T {
	if arg1 < arg2 {
		return arg2
	}
	return arg1
}

func (n *Node[Key, Data]) delete(key Key) *Node[Key, Data] {
	// node is nil -> nil
	if n == nil {
		return nil
	}

	if key < n.key {
		n.left = n.left.delete(key)
	} else if key > n.key {
		n.right = n.right.delete(key)
	} else {
		// delete this node case
		if n.right == nil && n.left == nil {
			// node has no child -> just delete
			n = nil
			return nil
		} else if n.right != nil && n.left != nil {
			// node has both children
			temp := n.right.minValueNode()
			n.key = temp.key
			n.right = n.right.delete(n.key)
		} else if n.right == nil && n.left != nil {
			// only left child
			n = n.left
		} else if n.right != nil && n.left == nil {
			// only right child
			n = n.right
		}
	}

	// node is nil -> nil
	if n == nil {
		fmt.Println("WUT??")
		return nil //! WTF check
	}

	if n.left != nil && n.right != nil {
		n.height = 1 + Max(n.left.height, n.right.height)
	} else if n.left == nil && n.right != nil {
		n.height = 1 + n.right.height
	} else if n.left != nil && n.right == nil {
		n.height = 1 + n.left.height
	} else {
		fmt.Println("KEK?")
	}

	return n.rebalance()
}

// func (n *Node[Key, Data]) delete(key Key) *Node[Key, Data] {
// 	if n == nil {
// 		return nil
// 	}
// 	if key < n.key {
// 		n.left = n.left.delete(key)
// 		return n.rebalance()
// 	}
// 	if key > n.key {
// 		n.right = n.right.delete(key)
// 		return n.rebalance()
// 	}

// 	if n.right == nil && n.left == nil {
// 		n = nil
// 		return nil
// 	}

// 	if n.left == nil {
// 		n = n.right
// 		return n.rebalance()
// 	}

// 	if n.right == nil {
// 		n = n.left
// 		return n.rebalance()
// 	}

// 	leftMostRightNode := n.right
// 	for {
// 		if leftMostRightNode != nil && leftMostRightNode.left != nil {
// 			leftMostRightNode = leftMostRightNode.left
// 		} else {
// 			break
// 		}
// 	}

// 	n.data = leftMostRightNode.data
// 	n.key = leftMostRightNode.key

// 	n.right = n.right.delete(n.key)

// 	n.height = max(n.left.getHeight(), n.right.getHeight()) + 1
// 	return n.rebalance()
// }

func (n *Node[Value, Data]) dump(i int, lr string) {
	if n == nil {
		return
	}
	indent := ""
	if i > 0 {
		indent = strings.Repeat(" ", (i-1)*4) + "+" + lr + "--"
	}
	fmt.Printf("%s%v[%d,%d]\n", indent, n.key, n.balanceFactor(), n.getHeight())
	n.left.dump(i+1, "L")
	n.right.dump(i+1, "R")
}

func max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
