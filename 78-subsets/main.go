package main

import (
	"fmt"
)

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) (res [][]int) {
	res = append(res, []int{})
	for _, n := range nums {
		for _, c := range res {
			kek := make([]int, len(c))
			copy(kek, c)
			c = append(kek, n)
			res = append(res, c)
		}
	}
	return
}

func subsetsRed(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res := make([][]int, 0)
	res = append(res, nums)
	for i := range nums {
		res = append(res, subsets(except(nums, i))...)
	}
	return res
}

func except(slice []int, i int) []int {
	a1, a2 := make([]int, i), make([]int, len(slice)-i-1)
	copy(a1, slice[:i])
	copy(a2, slice[i+1:])
	return append(a1, a2...)
}
