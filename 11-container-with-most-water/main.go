package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
}

func maxArea(height []int) int {
	maxVolume := 0
	for left, right := 0, len(height)-1; left < right; {
		volume := (right - left) * (min(height[left], height[right]))
		if volume > maxVolume {
			maxVolume = volume
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxVolume
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
