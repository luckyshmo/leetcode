package main

import "testing"

func BenchmarkXxx(b *testing.B) {
	m := [][]int{{50, 10, 90, 11}, {20, 40, 80, 10}, {13, 30, 60, 70}, {15, 14, 12, 16}}
	for i := 0; i < b.N; i++ {
		rotate(m)
	}
}
func BenchmarkKek(b *testing.B) {
	m := [][]int{{50, 10, 90, 11}, {20, 40, 80, 10}, {13, 30, 60, 70}, {15, 14, 12, 16}}
	for i := 0; i < b.N; i++ {
		rotateKek(m)
	}
}
