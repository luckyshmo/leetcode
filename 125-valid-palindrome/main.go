package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("0R"))
}

func isPalindrome(str string) bool {
	s := []rune(str)
	b, e := 0, len(s)-1
	for b < e {
		if !unicode.IsLetter(s[b]) && !unicode.IsNumber(s[b]) {
			b++
			continue
		}
		if !unicode.IsLetter(s[e]) && !unicode.IsNumber(s[e]) {
			e--
			continue
		}
		if unicode.ToLower(s[b]) != unicode.ToLower(s[e]) {
			// fmt.Printf("%c, %c", unicode.ToLower(s[b]), unicode.ToLower(s[e]))
			return false
		}
		b++
		e--
	}

	return true
}
