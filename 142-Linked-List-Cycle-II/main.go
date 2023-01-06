package main

import (
	"fmt"
)

func main() {
	fmt.Println()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(h *ListNode) *ListNode {
	slow, fast, head := h, h, h
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			for head != slow {
				head, slow = head.Next, slow.Next
			}
			return head
		}
	}
	return nil
}
