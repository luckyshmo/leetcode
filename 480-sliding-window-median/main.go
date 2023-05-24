package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(medianSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func medianSlidingWindow(nums []int, k int) []float64 {
	ans := make([]float64, 0)

	q := []int{}
	for i := 0; i < len(nums); i++ {
		if len(q) >= k {
			q = deleteNumber(q, nums[i-k])
		}
		q = insertNumber(q, nums[i])
		// fmt.Println(q)
		if len(q) == k && !even(k) {
			ans = append(ans, float64(q[k/2]))
		} else if len(q) == k {
			ans = append(ans, float64(q[k/2]+q[k/2-1])/2.0)
		}
	}

	return ans
}

func insertNumber(sortedArr []int, num int) []int {
	low, high := 0, len(sortedArr)-1

	if len(sortedArr) < 1 {
		return []int{num}
	}

	// Check if the number should be inserted at the beginning or end
	if num <= sortedArr[low] {
		return append([]int{num}, sortedArr...)
	} else if num >= sortedArr[high] {
		return append(sortedArr, num)
	}

	// Perform binary search to find the insertion position
	for low <= high {
		mid := (low + high) / 2

		if sortedArr[mid] == num {
			// Number already exists in the array, insert at mid position
			return append(sortedArr[:mid], append([]int{num}, sortedArr[mid:]...)...)
		} else if sortedArr[mid] < num {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	// Insert the number at the appropriate position
	insertIndex := low
	return append(sortedArr[:insertIndex], append([]int{num}, sortedArr[insertIndex:]...)...)
}

func deleteNumber(sortedArr []int, num int) []int {
	low, high := 0, len(sortedArr)-1

	for low <= high {
		mid := (low + high) / 2

		if sortedArr[mid] == num {
			// fmt.Println("before del: ", sortedArr)
			sortedArr = append(sortedArr[:mid], sortedArr[mid+1:]...)
			// fmt.Println("return after del: ", sortedArr)
			return sortedArr
		} else if sortedArr[mid] < num {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return sortedArr
}

func medianSlidingWindowSlow(nums []int, k int) []float64 {
	ans := make([]float64, 0)

	q := []int{}
	for i := 0; i < len(nums); i++ {
		if len(q) >= k {
			q = q[1:]
		}
		q = append(q, nums[i])
		a := make([]int, len(q))
		copy(a, q)
		sort.Ints(a)
		if len(a) == k && !even(len(a)) {
			ans = append(ans, float64(a[len(a)/2]))
		} else if len(a) == k {
			ans = append(ans, float64(a[len(a)/2]+a[len(a)/2-1])/2.0)
		}
	}

	return ans
}

func even(a int) bool {
	return (a ^ 1) == (a + 1)
}
