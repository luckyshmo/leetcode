package main

import (
	"fmt"
)

func main() {
	root := NewNode(Interval{15, 20}, 20)
	root.Insert(Interval{10, 30})
	root.Insert(Interval{17, 19})
	root.Insert(Interval{5, 20})
	root.Insert(Interval{12, 15})
	root.Insert(Interval{30, 40})

	root.inOrder()
}

type Node struct {
	r           Interval
	max         int
	left, right *Node
}

func NewNode(x Interval, max int) *Node {
	return &Node{
		r:   x,
		max: max,
	}
}

func (root *Node) inOrder() {
	if root == nil {
		return
	}
	root.left.inOrder()
	print()
	fmt.Println("interval=", root.r, "; max=", root.max)
	root.right.inOrder()
}

func (root *Node) isOverlapping(x Interval) (Interval, error) {
	if root == nil {
		// return a dummy interval range
		return [2]int{}, fmt.Errorf("nil root")
	}
	// if x overlaps with root's interval
	if (x.Low() > root.r.Low() &&
		x.Low() < root.r.High()) ||
		(x.High() > root.r.Low() &&
			x.High() < root.r.High()) {
		return root.r, nil
	} else if root.left != nil && root.left.max > x.Low() {
		// the overlapping node may be present in left child
		return root.left.isOverlapping(x)
	} else {
		return root.right.isOverlapping(x)
	}
}

func (root *Node) Insert(x Interval) *Node {
	if root == nil {
		return NewNode(x, x.High())
	}

	if x.Low() < root.r.Low() {
		root.left = root.left.Insert(x)
	} else {
		root.right = root.right.Insert(x)
	}

	if root.max < x.High() {
		root.max = x.High()
	}
	return root
}

type Interval [2]int

func (i Interval) High() int {
	return i[1]
}

func (i *Interval) Low() int {
	return i[0]
}
