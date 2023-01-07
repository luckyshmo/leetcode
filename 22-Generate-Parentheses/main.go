package main

func main() {
	// _, five := generateParenthesis(5)
	// exp := NewSet("((((()))))", "(((()())))", "(((())()))", "(((()))())", "(((())))()", "((()(())))", "((()()()))", "((()())())", "((()()))()", "((())(()))", "((())()())", "((())())()", "((()))(())", "((()))()()", "(()((())))", "(()(()()))", "(()(())())", "(()(()))()", "(()()(()))", "(()()()())", "(()()())()", "(()())(())", "(()())()()", "(())((()))", "(())(()())", "(())(())()", "(())()(())", "(())()()()", "()(((())))", "()((()()))", "()((())())", "()((()))()", "()(()(()))", "()(()()())", "()(()())()", "()(())(())", "()(())()()", "()()((()))", "()()(()())", "()()(())()", "()()()(())", "()()()()()")
	// fmt.Println(five.Diff(exp))
	generateParenthesis(4)
	// fmt.Println(combinations((7)))
}

func generateParenthesis(n int) (ans []string) {
	prev := make([]Set, 0)
	set := NewSet("()")
	for i := 2; i <= n; i++ {
		current := NewSet()
		if i > 3 {
			index := combinations(i)
			for _, ind := range index {
				c1, c2 := prev[ind[0]-2], prev[ind[1]-2]
				c1.Foreach(func(v string) {
					c2.Foreach(func(k string) {
						current.Add(v + k)
					})
				})
				c2.Foreach(func(v string) {
					c1.Foreach(func(k string) {
						current.Add(v + k)
					})
				})
			}
		}

		set.Foreach(func(op string) {
			a, b, c := op+"()", "()"+op, "("+op+")"
			current.Add(a, b, c)
		})
		prev = append(prev, *current)
		set = current
	}
	return set.Arr()
}

func combinations(n int) [][]int {
	var combinations [][]int
	for i := 2; i < n; i++ {
		j := n - i
		if i <= j {
			combinations = append(combinations, []int{i, j})
		}
	}
	return combinations
}

type Set struct {
	values map[string]struct{}
}

func NewSet(vals ...string) *Set {
	s := &Set{
		values: make(map[string]struct{}),
	}
	s.Add(vals...)
	return s
}

func (set *Set) Add(vals ...string) {
	for _, v := range vals {
		set.values[v] = struct{}{}
	}
}

func (set *Set) Remove(v string) {
	delete(set.values, v)
}

func (set *Set) Foreach(f func(v string)) {
	for v := range set.values {
		f(v)
	}
}

func (set *Set) Arr() (a []string) {
	for k := range set.values {
		a = append(a, k)
	}
	return
}

func (set *Set) Copy() *Set {
	newSet := NewSet()
	for k := range set.values {
		newSet.values[k] = struct{}{}
	}
	return newSet
}

func (set *Set) Diff(s *Set) (diff []string) {
	for k := range set.values {
		if _, ok := s.values[k]; !ok {
			diff = append(diff, k)
		}
	}
	for k := range s.values {
		if _, ok := set.values[k]; !ok {
			diff = append(diff, k)
		}
	}
	return
}
