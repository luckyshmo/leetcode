package main

import (
	"fmt"
)

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}

// XOR
func singleNumber(nums []int) (ans int) {
	for _, v := range nums {
		ans = ans ^ v
	}
	return ans
}

func singleNumberHashMap(nums []int) int {
	ans := make(map[int]struct{}, 0)
	for _, val := range nums {
		if _, ok := ans[val]; ok {
			delete(ans, val)
		} else {
			ans[val] = struct{}{}
		}
	}

	for k := range ans {
		return k
	}
	return -1
}
