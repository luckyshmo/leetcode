package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println(mergeKLists(nil))
}

func mergeKLists(lists []*ListNode) *ListNode {
	res := &ListNode{-1, nil}
	merged := res
	pq := PQueue{}
	for _, l := range lists {
		if l != nil {
			pq = append(pq, l)
		}
	}
	heap.Init(&pq)
	for len(pq) > 0 {
		merged.Next = heap.Pop(&pq).(*ListNode)
		merged = merged.Next
		if merged.Next != nil {
			heap.Push(&pq, merged.Next)
		}
	}
	return res.Next
}

type PQueue []*ListNode

func (pq PQueue) Len() int {
	return len(pq)
}

func (pq PQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQueue) Push(x interface{}) {
	item := x.(*ListNode)
	*pq = append(*pq, item)
}

func (pq *PQueue) Pop() interface{} {
	old := *pq
	item := old[len(*pq)-1]
	*pq = old[:len(*pq)-1]
	return item
}

func (pq *PQueue) Peek(x interface{}) *ListNode {
	return (*pq)[len(*pq)-1]
}
