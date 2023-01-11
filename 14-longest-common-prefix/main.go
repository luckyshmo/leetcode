package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"123", "123as", "123asd"}))
}

func longestCommonPrefix(strs []string) (common string) {
	for i := range strs[0] {
		if ok := isTheSame(strs, i); ok {
			common += string(strs[0][i])
		} else {
			return
		}
	}
	return
}

func isTheSame(strs []string, i int) (isTheSame bool) {
	for j := 1; j <= len(strs)-1; j++ {
		if len(strs[j]) <= i {
			return
		}
		if strs[0][i] != strs[j][i] {
			return
		}
	}
	return true
}
