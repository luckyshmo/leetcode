package main

import (
	"fmt"
)

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("abcde", "edcba"))
	fmt.Println(isAnagram("ab", "a"))
}

func isAnagram(s string, t string) bool {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	for _, v := range t {
		if val, ok := m[v]; ok {
			if val == 1 {
				delete(m, v)
				continue
			}
			m[v]--
			continue
		}
		return false
	}

	return len(m) == 0
}
