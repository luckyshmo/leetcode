package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println()
}

func topKFrequent(nums []int, k int) []int {
	pq := PriorityQueue{}
	m := make(map[int]int, 0) // number -> count
	for _, n := range nums {
		m[n]++
	}
	for value, count := range m {
		pq = append(pq, &Item{
			value: value,
			count: count,
		})
	}
	heap.Init(&pq)

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(&pq).(*Item).value
	}

	return res
}

// An Item is something we manage in a priority queue.
type Item struct {
	value int
	count int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].count > pq[j].count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
