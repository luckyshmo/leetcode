package main

import "fmt"

func main() {
	r1 := permute([]int{6, 3, 2, 7, 4, -1})
	r2 := permute([]int{6, 3, 2, 7, 4, -1})

	diff := make([][]int, 0)

	for _, arr1 := range r1 {
		isUnique := true
		for _, arr2 := range r2 {
			if Equal(arr1, arr2) {
				isUnique = false
			}
		}
		if isUnique {
			diff = append(diff, arr1)
		}
	}
	fmt.Println(diff)

	for _, arr2 := range r1 {
		isUnique := true
		for _, arr1 := range r2 {
			if Equal(arr1, arr2) {
				isUnique = false
			}
		}
		if isUnique {
			diff = append(diff, arr2)
		}
	}
	fmt.Println(diff)

	// fmt.Println(len(permute([]int{6, 3, 2, 7, 4, -1})))
	// fmt.Println(len(permuteKek([]int{6, 3, 2, 7, 4, -1})))
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func permute(nums []int) (p [][]int) {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	for j := range nums {
		withoutJ := append(nums[j+1:], nums[:j]...)
		c := permute(withoutJ)
		for i := range c {
			c[i] = append(c[i], nums[j])
		}
		p = append(p, c...)
	}

	return p
}

func permuteKek(nums []int) [][]int {
	var output [][]int
	var backtrack func(first int)
	backtrack = func(first int) {
		if first == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, nums)
			output = append(output, temp)
			return
		}
		for i := first; i < len(nums); i++ {
			nums[first], nums[i] = nums[i], nums[first]
			backtrack(first + 1)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}
	backtrack(0)
	return output
}
