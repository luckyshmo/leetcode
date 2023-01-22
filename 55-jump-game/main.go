package main

import (
	"fmt"
)

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}

func canJump(nums []int) bool {
	maxIndex := len(nums) - 1
	max := 0
	for i, num := range nums {
		cmax := i + num
		if cmax >= maxIndex {
			return true
		}
		if cmax > max {
			max = cmax
		}
		if max == i {
			return false
		}
	}
	return false
}
