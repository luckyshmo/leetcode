package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{-3, -1, 0, 0, 0, 1, 2, 4, 5}, 1))
}

func fourSum(n []int, target int) [][]int {
	res := make(map[[4]int]struct{}, 0)
	l := len(n)
	kek := 0
	sort.Ints(n)
	a, b, c, d := 0, 1, 2, 3
	for a < l-3 && a < b {
		b = a + 1
		for b < l-2 && b < c {
			c = b + 1
			for c < l-1 && c < d {
				d = c + 1
				for d < l {
					sum := n[a] + n[b] + n[c] + n[d]
					if sum == target {
						kek++
						res[[4]int{n[a], n[b], n[c], n[d]}] = struct{}{}
					} else if sum > target {
						d++
						continue
					}
					d++
				}
				c++
			}
			b++
		}
		a++
	}
	result := make([][]int, 0)
	for k := range res {
		r := make([]int, 4)
		copy(r, k[:])
		result = append(result, r)
	}
	return result
}
