package main

func main() {
	println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}

func search(nums []int, target int) int {
	index := 0
	for len(nums) > 1 {
		mi := len(nums) / 2
		m := nums[mi]
		if m < target {
			nums = nums[mi:]
			index += mi
		} else if m > target {
			nums = nums[:mi]
		} else {
			return index + mi
		}
	}
	if nums[0] == target {
		return index
	}
	return -1
}

// func search(nums []int, target int) int {
// 	return rSearch(nums, target, 0)
// }

func rSearch(nums []int, target, i int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return i
		}
		return -1
	}
	mi := len(nums) / 2
	m := nums[mi]
	if m < target {
		return rSearch(nums[mi:], target, i+mi)
	} else if m > target {
		return rSearch(nums[:mi], target, i)
	} else {
		return i + mi
	}
}
