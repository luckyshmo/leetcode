package main

import (
	"fmt"
)

var a = 1

func main() {
	fmt.Println(guessNumber(2))
}

func guessNumber(n int) int {
	return g(1, n)
}

func g(b, e int) int {
	if e-b < 2 {
		if a := guess(e); a == 0 {
			return e
		}
		return b
	}
	m := b + (e-b)/2
	switch res := guess(m); res {
	case -1:
		return g(b, m)
	case 0:
		return m
	case 1:
		return g(m, e)
	}
	return -1
}

func guess(g int) int {
	if g == a {
		return 0
	} else if g < a {
		return 1
	} else {
		return -1
	}
}
