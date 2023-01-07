package main

type Set[T comparable] struct {
	values map[T]struct{}
}

func NewSet[T comparable](vals ...T) *Set[T] {
	s := &Set[T]{
		values: make(map[T]struct{}),
	}
	s.Add(vals...)
	return s
}

func (set *Set[T]) Add(vals ...T) {
	for _, v := range vals {
		set.values[v] = struct{}{}
	}
}

func (set *Set[T]) Remove(v T) {
	delete(set.values, v)
}

func (set *Set[T]) Foreach(f func(v T)) {
	for v := range set.values {
		f(v)
	}
}

func (set *Set[T]) Arr() (a []T) {
	for k := range set.values {
		a = append(a, k)
	}
	return
}

func (set *Set[T]) Copy() *Set[T] {
	newSet := NewSet[T]()
	for k := range set.values {
		newSet.values[k] = struct{}{}
	}
	return newSet
}
