package main

import (
	"container/heap"
	"fmt"
)

func main() {
	kek := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(kek.Add(3))
	fmt.Println(kek.Add(5))
	fmt.Println(kek.Add(10))
	fmt.Println(kek.Add(9))
	fmt.Println(kek.Add(4))
}

type KthLargest struct {
	k  int
	pq *PQueue
}

func Constructor(k int, nums []int) KthLargest {
	pq := PQueue{}
	heap.Init(&pq)
	for _, n := range nums {
		heap.Push(&pq, n)
	}

	// remove the smaller elements from the heap such that only the k largest elements are in the heap
	for len(pq) > k {
		heap.Pop(&pq)
	}
	return KthLargest{
		k:  k,
		pq: &pq,
	}
}

func (this *KthLargest) Add(val int) int {
	if len(*this.pq) < this.k {
		heap.Push(this.pq, val)
	} else if val > (*this.pq)[0] {
		heap.Push(this.pq, val)
		heap.Pop(this.pq)
	}
	return (*this.pq)[0]
}

type PQueue []int

func (pq PQueue) Len() int {
	return len(pq)
}

func (pq PQueue) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq PQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
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
