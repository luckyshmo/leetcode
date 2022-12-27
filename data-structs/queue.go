package main

type Node[T any] struct {
	Val      T
	previous *Node[T]
	next     *Node[T]
}

type Queue[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func (q *Queue[T]) Push(val T) {
	newNode := &Node[T]{
		Val: val,
	}
	// Empty queue
	if q.head == nil && q.tail == nil {
		q.head = newNode
		q.tail = newNode
	}
	// Non empty queue
	q.head.next = newNode
	temp := q.head
	q.head = newNode
	q.head.previous = temp
}

func (q *Queue[T]) Pop() (val T) {
	// empty queue
	if q.tail == nil {
		return
	}
	res := q.tail.Val
	q.tail = q.tail.next
	return res
}
