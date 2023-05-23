package main

import (
	"fmt"
)

func main() {
	fmt.Println()
}

func partitionLabels(s string) []int {
	last := make(map[rune]int)
	for i, c := range s {
		last[c] = i
	}
	j := 0
	anchor := 0
	ans := []int{}
	for i, c := range s {
		j = max(j, last[c])
		if i == j {
			ans = append(ans, i-anchor+1)
			anchor = i + 1
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
