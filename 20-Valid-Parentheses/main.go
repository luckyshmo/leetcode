package main

import (
	"fmt"
)

func main() {
	isValid(")")
	fmt.Println(isValid("()"))
	fmt.Println(isValid("({[]})"))
	fmt.Println(isValid("(("))
	fmt.Println(isValid("(}"))
	fmt.Println(isValid("((){)}"))
}

var closeToOpen = map[rune]rune{
	rune(")"[0]): rune("("[0]),
	rune("}"[0]): rune("{"[0]),
	rune("]"[0]): rune("["[0]),
}

func isValid(s string) bool {
	st := NewStack()
	for _, c := range s {
		v, ok := closeToOpen[c]
		if !ok {
			st.Push(c)
			continue
		}
		head := st.Pop()
		if head != v {
			return false
		}
	}
	st.Print()
	return st.IsEmpty()
}

type node struct {
	previous *node
	val      rune
}

type Stack struct {
	head *node
}

func NewStack(vals ...rune) *Stack {
	var s Stack
	for _, v := range vals {
		s.Push(v)
	}
	return &s
}

func (s *Stack) Push(val rune) {
	s.head = &node{
		val:      val,
		previous: s.head,
	}
}

func (s *Stack) Pop() (val rune) {
	if s.head == nil {
		return -1
	}
	headVal := s.head.val
	s.head = s.head.previous
	return headVal
}

func (s *Stack) Peek() (val rune) {
	return s.head.val
}

func (s *Stack) IsEmpty() bool {
	return s.head == nil
}

func (s *Stack) Print() {
	for n := s.head; n != nil; n = n.previous {
		fmt.Print(n.val, ", ")
	}
	fmt.Println()
}
