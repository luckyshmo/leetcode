package main

import "testing"

func BenchmarkKek(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateParenthesisKek(8)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateParenthesis(8)
	}
}
