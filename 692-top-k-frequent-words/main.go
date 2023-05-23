package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println()
}

func topKFrequent(words []string, k int) []string {
	pq := PriorityQueue{}
	m := make(map[string]int, 0) // number -> count
	for _, w := range words {
		m[w]++
	}
	for value, count := range m {
		pq = append(pq, &Item{
			value: value,
			count: count,
		})
	}
	heap.Init(&pq)

	res := make([]string, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(&pq).(*Item).value
	}

	return res
}

// An Item is something we manage in a priority queue.
type Item struct {
	value string
	count int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].count == pq[j].count {
		return pq[i].value < pq[j].value
	}
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
