package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(myAtoi("  0000000000012345678"))
	fmt.Println(myAtoi("00000-42a1234"))
	fmt.Println(myAtoi("-42a1234"))
	fmt.Println(myAtoi("+"))
	fmt.Println(myAtoi("-"))
	fmt.Println(myAtoi("-1"))
	fmt.Println(myAtoi("1"))
	fmt.Println(myAtoi(""))
	fmt.Println(myAtoi(" "))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("  -42"))
	fmt.Println(myAtoi("  -42 asdasd"))
	fmt.Println(myAtoi("+-12"))
}

func myAtoi(s string) int {
	if s == "" {
		return 0
	}
	zero := []rune("0")[0]
	space := []rune(" ")[0]
	minus := []rune("-")[0]
	plus := []rune("+")[0]

	numbers := map[rune]struct{}{
		zero:           {},
		[]rune("1")[0]: {},
		[]rune("2")[0]: {},
		[]rune("3")[0]: {},
		[]rune("4")[0]: {},
		[]rune("5")[0]: {},
		[]rune("6")[0]: {},
		[]rune("7")[0]: {},
		[]rune("8")[0]: {},
		[]rune("9")[0]: {},
	}

	var start, end, i int
	var isZero bool
	var r rune
	for i, r = range s {
		if r == space {
			if isZero {
				return 0
			}
			continue
		}
		if r == zero {
			isZero = true
			continue
		}
		if r == minus {
			if isZero {
				return 0
			}
			start = i
			i++
			break
		}
		if r == plus {
			if isZero {
				return 0
			}
			i++
			start = i
			break
		}
		_, isNumber := numbers[r]
		if !isNumber {
			return 0
		}
		start = i
		break
	}

	for ; i < len(s); i++ {
		_, isNumber := numbers[rune(s[i])]
		if isNumber {
			end = i
			continue
		}
		i++
		break
	}

	if start > end {
		return 0
	}
	str := s[start : end+1]
	s1, _ := strconv.Atoi(str)
	if s1 > 2147483647 {
		return 2147483647
	}
	if s1 < -2147483648 {
		return -2147483648
	}
	return s1
}
