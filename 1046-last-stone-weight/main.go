package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
}

func lastStoneWeight(stones []int) int {
	pq := PQueue{}
	for _, s := range stones {
		pq = append(pq, s)
	}
	heap.Init(&pq)
	for len(pq) >= 2 {
		st1, st2 := heap.Pop(&pq).(int), heap.Pop(&pq).(int)
		m := mod(st1, st2)
		if m > 0 {
			heap.Push(&pq, m)
		}
	}
	if pq.Len() == 1 {
		return heap.Pop(&pq).(int)
	}
	return 0
}

func mod(a, b int) int {
	m := a - b
	if m < 0 {
		return -m
	}
	return m
}

type PQueue []int

func (pq PQueue) Len() int {
	return len(pq)
}

func (pq PQueue) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq PQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQueue) Push(x interface{}) {
	item := x.(int)
	*pq = append(*pq, item)
}

func (pq *PQueue) Pop() interface{} {
	old := *pq
	item := old[len(*pq)-1]
	*pq = old[:len(*pq)-1]
	return item
}

func (pq PQueue) Peek(x interface{}) int {
	return (pq)[len(pq)-1]
}
