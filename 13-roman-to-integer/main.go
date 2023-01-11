package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) (num int) {
	var next byte = 0
	for i := len(s) - 1; i >= 0; i-- {
		r := s[i]
		switch r {
		case 'I':
			if next == 'V' || next == 'X' {
				num--
				continue
			}
			num++
		case 'V':
			num += 5
		case 'X':
			if next == 'L' || next == 'C' {
				num -= 10
				continue
			}
			num += 10
		case 'L':
			num += 50
		case 'C':
			if next == 'D' || next == 'M' {
				num -= 100
				continue
			}
			num += 100
		case 'D':
			num += 500
		case 'M':
			num += 1000
		}
		next = r
	}
	return
}
