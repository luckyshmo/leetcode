package main

import (
	"testing"
)

func Benchmark_Range(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxProfit([]int{7, 1, 5, 3, 6, 4})
	}
}

func Benchmark_Iter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxProfitRange([]int{7, 1, 5, 3, 6, 4})
	}
}
