package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) print() {
	for ; l != nil; l = l.Next {
		fmt.Print(l.Val, ", ")
	}
	fmt.Println()
}

func main() {
	n5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	n4 := &ListNode{
		Val:  4,
		Next: n5,
	}
	n3 := &ListNode{
		Val:  3,
		Next: n4,
	}
	n2 := &ListNode{
		Val:  2,
		Next: n3,
	}
	n1 := &ListNode{
		Val:  1,
		Next: n2,
	}
	n1.print()
	ans := removeNthFromEnd(n1, 1)
	ans.print()

	mek := &ListNode{
		Val: 2,
	}
	kek := &ListNode{
		Val:  1,
		Next: mek,
	}
	lol := removeNthFromEnd(kek, 1)
	lol.print()
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p2 := head
	var p1 *ListNode
	for i := 0; p2 != nil; p2, i = p2.Next, i+1 {
		if i == n {
			p1 = head
		} else if i > n {
			p1 = p1.Next
		}
		if p2.Next == nil {
			if p1 == nil {
				return head.Next
			}
			p1.Next = p1.Next.Next
		}
	}
	return head
}
