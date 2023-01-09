package main

import (
	"fmt"
)

func main() {
	n1 := &ListNode{
		Val: 3,
	}
	n2 := &ListNode{
		Val:  2,
		Next: n1,
	}
	n3 := &ListNode{
		Val:  1,
		Next: n2,
	}
	k1 := &ListNode{
		Val: 3,
	}
	k2 := &ListNode{
		Val:  2,
		Next: k1,
	}
	k3 := &ListNode{
		Val:  1,
		Next: k2,
	}
	mergeTwoLists(n3, k3).Print()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {
	for ; l != nil; l = l.Next {
		fmt.Println(l.Val, ", ")
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val: -1,
	}
	kek := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = &ListNode{
				Val: l1.Val,
			}
			l1 = l1.Next
		} else {
			head.Next = &ListNode{
				Val: l2.Val,
			}
			l2 = l2.Next
		}
		head = head.Next
	}
	if l1 != nil {
		head.Next = l1
	}
	if l2 != nil {
		head.Next = l2
	}
	return kek.Next
}
