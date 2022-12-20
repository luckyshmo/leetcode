package main

import "fmt"

func main() {
	// s := "babad"
	s := "ac"
	fmt.Print(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	longest := ""
	for i := 0; i < len(s)-1; i++ {
		for j := len(s) - 1; j >= i; j-- {
			if s[i] == s[j] {
				mb := s[i : j+1]
				p := IsPalindrome(mb)
				if p && len(longest) < len(mb) {
					longest = s[i : j+1]
				}
			}
		}
	}
	return longest
}

func IsPalindrome(p string) bool {
	for i, j := 0, len(p)-1; i < j; i, j = i+1, j-1 {
		if p[i] != p[j] {
			return false
		}
	}
	return true
}
