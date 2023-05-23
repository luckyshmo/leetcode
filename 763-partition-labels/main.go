package main

import (
	"fmt"
)

func main() {
	fmt.Println()
}

func partitionLabels(S string) []int {
	last := make(map[rune]int)
	for i, c := range S {
		last[c] = i
	}
	j := 0
	anchor := 0
	ans := []int{}
	for i, c := range S {
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
