package main

import (
	"fmt"
)

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	fmt.Println(rob([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
}

var cache []int

func rob(n []int) int {
	cache = make([]int, len(n)+1)
	for i := 0; i < len(cache); i++ {
		cache[i] = -1
	}
	return robHelper(n)
}

func robHelper(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	if cache[len(nums)] != -1 {
		return cache[len(nums)]
	}
	cache[len(nums)] = max(nums[0]+robHelper(nums[2:]), robHelper(nums[1:]))
	return cache[len(nums)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
