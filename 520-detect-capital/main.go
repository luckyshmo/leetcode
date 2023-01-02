package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "ǳ"
	fmt.Println(strings.ToTitle(str))
	fmt.Println(strings.ToUpper(str))
	fmt.Println(detectCapitalUse("Leetcode"))
}

func detectCapitalUse(word string) bool {
	upper := strings.ToUpper(word)
	lower := strings.ToLower(word)
	title := word[:1] + strings.ToLower(word[1:])
	if word == upper || word == lower || title == word {
		return true
	}
	return false
}
