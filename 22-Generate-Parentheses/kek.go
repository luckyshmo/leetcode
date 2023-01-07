package main

var ps []string
var cache []byte

func generateParenthesisKek(n int) []string {
	ps = make([]string, 0)
	cache = make([]byte, 0, n)
	backtrack(n, n, n)
	return ps
}

func backtrack(n, open, close int) {
	if close == 0 {
		ps = append(ps, string(cache))
		return
	}

	if open > 0 {
		cache = append(cache, '(')
		backtrack(n, open-1, close)
		cache = cache[:len(cache)-1]
	}

	if close > open {
		cache = append(cache, ')')
		backtrack(n, open, close-1)
		cache = cache[:len(cache)-1]
	}
}
