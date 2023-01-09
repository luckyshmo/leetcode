package main

import (
	"fmt"
)

func main() {
	fmt.Println(strStr("12abc12", "abc"))
	fmt.Println(strStr("mississippi", "issip"))
}

func strStr(str string, ss string) int {
	for i := 0; i < len(str); i++ {
		for j := 0; i+j < len(str) && str[i+j] == ss[j]; j++ {
			if j == len(ss)-1 {
				return i
			}
		}
	}
	return -1
}
