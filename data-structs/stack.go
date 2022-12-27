package main

import (
	"fmt"
)

type node[T any] struct {
	previous *node[T]
	val      T
}

type Stack[T any] struct {
	head *node[T]
}

func NewStack[T any](vals ...T) *Stack[T] {
	var s Stack[T]
	for _, v := range vals {
		s.Push(v)
	}
	return &s
}

func (s *Stack[T]) Push(val T) {
	s.head = &node[T]{
		val:      val,
		previous: s.head,
	}
}

func (s *Stack[T]) Pop() (val T) {
	headVal := s.head.val
	s.head = s.head.previous
	return headVal
}

func (s *Stack[T]) Peek() (val T) {
	return s.head.val
}

func (s *Stack[T]) IsEmpty() bool {
	return s.head == nil
}

func (s *Stack[T]) Print() {
	for n := s.head; n != nil; n = n.previous {
		fmt.Print(n.val, ", ")
	}
	fmt.Println()
}
