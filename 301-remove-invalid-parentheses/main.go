package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeInvalidParentheses("()())()"))
	// fmt.Println(removeInvalidParentheses("a("))
}

func removeInvalidParentheses(s string) []string {
	set = map[string]bool{}
	ans = make([][]string, len(s))
	rec(s)
	for i := len(s) - 1; i >= 0; i-- {
		if ans[i] != nil {
			return ans[i]
		}
	}
	return []string{""}
}

var (
	set map[string]bool
	ans [][]string
)

var parentheses1 = map[byte]struct{}{
	")"[0]: {},
	"("[0]: {},
	"}"[0]: {},
	"{"[0]: {},
	"]"[0]: {},
	"["[0]: {},
}

func rec(s string) {
	if s == "" {
		return
	}
	if _, ok := set[s]; ok {
		return
	}
	set[s] = isValid(s)
	if set[s] {
		ans[len(s)-1] = append(ans[len(s)-1], s)
		return // can do so bcs where is no reason to search subsets
	}
	for i := 0; i < len(s); i++ {
		if _, ok := parentheses1[s[i : i+1][0]]; ok {
			rec(s[:i] + s[i+1:])
		}
	}
}

var closeToOpen = map[rune]rune{
	rune(")"[0]): rune("("[0]),
	rune("}"[0]): rune("{"[0]),
	rune("]"[0]): rune("["[0]),
}

func isValid(s string) bool {
	st := NewStack()
	for _, c := range s {
		if _, ok := parentheses1[byte(c)]; ok {
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
	}
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
