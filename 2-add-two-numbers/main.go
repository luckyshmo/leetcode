package main

import "fmt"

func main() {
	l1 := ListNode{
		Val: 9,
		// Next: &ListNode{
		// 	Val: 9,
		// 	Next: &ListNode{
		// 		Val:  9,
		// 		Next: nil,
		// 	},
		// },
	}
	l2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 9,
				},
			},
		},
	}
	addTwoNumbers(&l1, &l2).Print()
}

func (ln *ListNode) Print() {
	for currentNode := ln; currentNode != nil; currentNode = currentNode.Next {
		fmt.Print(currentNode.Val, ",")
	}
	fmt.Println()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddRemainder(ln *ListNode, r int) {
	firstDigit, lastDigit := r, 0
	for currentNode := ln; firstDigit > 0; currentNode = currentNode.Next {
		lastDigit, firstDigit = Sum(currentNode.Val, firstDigit, 0)
		if currentNode.Next == nil && currentNode.Val+firstDigit >= 10 {
			currentNode.Val = lastDigit
			currentNode.Next = &ListNode{
				Val: firstDigit,
			}
			break
		} else if currentNode.Next == nil {
			currentNode.Val = lastDigit + firstDigit
			break
		}
		currentNode.Val = lastDigit
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum, remainder := Sum(l1.Val, l2.Val, 0)
	ans := &ListNode{
		Val:  sum,
		Next: nil,
	}

	for currentNode, tl1, tl2 := ans, l1.Next, l2.Next; ; currentNode, tl1, tl2 = currentNode.Next, tl1.Next, tl2.Next {
		if tl2 == nil && tl1 == nil {
			if remainder > 0 {
				currentNode.Next = &ListNode{
					Val:  remainder,
					Next: nil,
				}
			}
			break
		} else if tl1 == nil {
			currentNode.Next = tl2
			if remainder > 0 {
				AddRemainder(currentNode.Next, remainder)
			}
			break
		} else if tl2 == nil {
			currentNode.Next = tl1
			if remainder > 0 {
				AddRemainder(currentNode.Next, remainder)
			}
			break
		}
		//------
		sum, newRemainder := Sum(tl1.Val, tl2.Val, remainder)
		newNode := ListNode{
			Val:  sum,
			Next: nil,
		}
		currentNode.Next = &newNode
		remainder = newRemainder
	}

	return ans
}

func Sum(a, b, r int) (lastDigit, firstDigit int) {
	sum := a + b + r
	return sum % 10, sum / 10
}
