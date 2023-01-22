package main

import (
	"fmt"
)

func main() {
	fmt.Println(spiralOrder([][]int{{2, 5, 8}, {4, 0, -1}}))
	// fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}

func spiralOrder(matrix [][]int) []int {
	xLen := len(matrix[0])
	yLen := len(matrix)
	x, y, step := 0, 0, 0
	res := make([]int, 0)
	for step < xLen/2+1 && step < yLen/2+1 && len(res) < xLen*yLen {
		res = append(res, matrix[step][step:xLen-step]...)
		x = xLen - step - 1
		for y = step + 1; y < yLen-step && len(res) < xLen*yLen; y++ {
			res = append(res, matrix[y][x])
		}
		y--
		x--
		for ; x >= step && len(res) < xLen*yLen; x-- {
			res = append(res, matrix[y][x])
		}
		x++
		y--
		for ; y > step && len(res) < xLen*yLen; y-- {
			res = append(res, matrix[y][x])
		}
		step++
	}

	return res
}
