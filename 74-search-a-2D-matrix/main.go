package main

import (
	"fmt"
)

func main() {
	fmt.Println(searchMatrix([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}, 5))
}

func searchMatrix(matrix [][]int, target int) bool {
	x, y := len(matrix[0]), len(matrix)
	left, right := 0, x*y-1

	for left != right {
		mid := (left + right - 1) / 2
		if matrix[mid/x][mid%x] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return matrix[right/x][right%x] == target
}
