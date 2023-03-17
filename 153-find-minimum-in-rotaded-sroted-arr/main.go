package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMin([]int{4, 5, 1, 2, 3}))
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{11, 13, 15, 17}))
	fmt.Println(findMin([]int{2, 1}))
	fmt.Println(findMin([]int{3, 1, 2}))
}

func findMin(nums []int) int {
	if len(nums) <= 3 {
		return min(nums...)
	}
	for len(nums) > 3 {
		mi := len(nums) / 2
		first, middle, last := nums[0], nums[mi], nums[len(nums)-1]

		fIsM := first <= middle
		sIsM := last >= middle

		if !sIsM {
			nums = nums[mi+1:]
		} else if !fIsM {
			nums = nums[:mi+1]
		} else { // sIsM && fIsM
			return nums[0] // return zero elem of monotonic sequence + index
		}
	}
	return min(nums...)
}

func min(nums ...int) int {
	if len(nums) == 1 {
		return nums[0]
	} else {
		m := len(nums) / 2
		a := min(nums[:m]...)
		b := min(nums[m:]...)
		if a > b {
			return b
		}
		return a
	}
}
