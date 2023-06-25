package main

import (
	"fmt"
)

func main() {
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
	// fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}

func summaryRanges(nums []int) []string {
	res := make([]string, 0)
	for i := 0; i < len(nums); i++ {
		s := fmt.Sprint(nums[i])

		for ; i < len(nums)-1 && nums[i+1]-1 == nums[i]; i++ {
		}

		if s != fmt.Sprint(nums[i]) {
			s = s + "->" + fmt.Sprint(nums[i])
		}

		res = append(res, s)
	}

	return res
}
