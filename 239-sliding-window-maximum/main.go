package main

import (
	"fmt"
)

func main() {
	fmt.Println()
}

func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0)

	q := queue{make([]int, 0)}
	for i := 0; i < len(nums); i++ {
		if !q.isEmpty() && q.first() < i-k+1 {
			q.dequeue()
		}
		for !q.isEmpty() && nums[i] >= nums[q.last()] {
			q.pop()
		}
		q.enqueue(i)
		if i >= k-1 {
			ans = append(ans, nums[q.first()])
		}
	}

	return ans
}

type queue struct {
	a []int
}

func (q *queue) isEmpty() bool {
	return len(q.a) == 0
}

func (q *queue) enqueue(v int) {
	q.a = append(q.a, v)
}

func (q *queue) dequeue() {
	q.a = q.a[1:]
}

func (q *queue) pop() {
	q.a = q.a[:len(q.a)-1]
}

func (q *queue) first() int {
	return q.a[0]
}

func (q *queue) last() int {
	return q.a[len(q.a)-1]
}
