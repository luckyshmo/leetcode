package main

import (
	"fmt"
)

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 5))
}

func search(nums []int, target int) int {
	base := 0
	for len(nums) != 0 {
		middleInd := len(nums) / 2
		first, middle, last := nums[0], nums[middleInd], nums[len(nums)-1]
		if target == first {
			return base
		}
		if target == middle {
			return base + middleInd
		}
		if target == last {
			return base + len(nums) - 1
		}

		firstHalfIsMonotonic := first < middle
		secondHalfIsMonotonic := last > middle
		targetInFirstHalf := target > first && target < middle
		targetInSecondHalf := target > middle && target < last

		if (firstHalfIsMonotonic && targetInFirstHalf) || (secondHalfIsMonotonic && !targetInSecondHalf) {
			nums = nums[:middleInd]
		} else if (secondHalfIsMonotonic && targetInSecondHalf) || (firstHalfIsMonotonic && !targetInFirstHalf) {
			nums = nums[middleInd+1:]
			base += middleInd + 1
		} else {
			return -1
		}
	}
	return -1
}
