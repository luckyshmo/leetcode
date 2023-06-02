package main

import (
	"fmt"
)

func main() {
	fmt.Println()
}

func maxDistToClosest(s []int) (max int) {
	for i := 0; i < len(s); i++ {
		c := 0
		for ; i < len(s) && s[i] == 0; i, c = i+1, c+1 {
		}
		if !(c == i || i == len(s)) {
			if c%2 == 0 {
				c = c / 2
			} else {
				c = c/2 + 1
			}
		}
		if c > max {
			max = c
		}
	}

	return
}
