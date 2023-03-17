package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n5 := ListNode{
		Val:  5,
		Next: nil,
	}
	n4 := ListNode{
		Val:  4,
		Next: &n5,
	}
	n3 := ListNode{
		Val:  3,
		Next: &n4,
	}
	n2 := ListNode{
		Val:  2,
		Next: &n3,
	}
	n1 := ListNode{
		Val:  1,
		Next: &n2,
	}
	ans := reverseList(&n1)
	fmt.Println(ans)
}

// func reverseList(head *ListNode) *ListNode {
// 	var next, previous *ListNode
// 	for head != nil {
// 		next = head.Next
// 		head.Next = previous
// 		previous = head
// 		head = next
// 	}

// 	return previous
// }

func reverseList(head *ListNode) *ListNode {
	return reverseReq(nil, head)
}

func reverseReq(prev, head *ListNode) *ListNode {
	if head == nil {
		return prev
	}
	next := head.Next
	head.Next = prev
	prev = head
	head = next

	return reverseReq(prev, head)
}
