package main

import "fmt"

func main() {
	res := []int{1, 2, 3, 0, 0, 0}
	merge(res, 3, []int{2, 5, 6}, 3)
	fmt.Println(res)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	merged := make([]int, m+n)
	i1, i2 := 0, 0
	for i1 < m && i2 < n {
		if nums1[i1] < nums2[i2] {
			merged[i1+i2] = nums1[i1]
			i1++
		} else {
			merged[i1+i2] = nums2[i2]
			i2++
		}
	}
	for ; i1 < m; i1++ {
		merged[i1+i2] = nums1[i1]
	}
	for ; i2 < n; i2++ {
		merged[i1+i2] = nums2[i2]
	}

	copy(nums1, merged)
}
