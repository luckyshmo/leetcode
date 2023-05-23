package main

import (
	"fmt"
)

func main() {
	fmt.Println(characterReplacement("BBABAAB", 1))
	fmt.Println(characterReplacement("AABABBA", 1))
}

func characterReplacement(s string, k int) int {
	len := len(s)
	count := make([]int, 26)
	start, maxCount, maxLength := 0, 0, 0
	for end := 0; end < len; end++ {
		count[s[end]-'A']++
		maxCount = max(maxCount, count[s[end]-'A'])
		for end-start+1-maxCount > k {
			count[s[start]-'A']--
			start++
		}
		maxLength = max(maxLength, end-start+1)
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
