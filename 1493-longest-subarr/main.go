package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}))
}

func longestSubarray(nums []int) int {
	l, r := 0, 0
	dPos := -1
	max := 0
	for ; r < len(nums); r++ {
		if nums[r] == 0 && dPos == -1 {
			dPos = r
		} else if nums[r] == 0 && dPos != -1 {
			l = dPos + 1
			dPos = r
		}
		if max < r-l {
			max = r - l
		}
	}

	return max
}
