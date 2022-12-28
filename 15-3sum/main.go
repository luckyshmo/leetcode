package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}))
	fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
}

func threeSum(nums []int) [][]int {
	ans := [][]int{}
	ans1 := make(map[string]struct{})

	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				key := fmt.Sprint(nums[i], nums[j], nums[k])
				_, isDouble := ans1[key]
				if !isDouble {
					ans = append(ans, []int{nums[i], nums[j], nums[k]})
					ans1[key] = struct{}{}
					continue
				}
				j++
				k--
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return ans
}
