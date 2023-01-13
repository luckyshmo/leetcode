package main

import (
	"fmt"
)

func main() {
	m := [][]int{{50, 10, 90, 11}, {20, 40, 80, 10}, {13, 30, 60, 70}, {15, 14, 12, 16}}
	print(m)
	rotate(m)
	print(m)
}

func print(matrix [][]int) {
	for _, line := range matrix {
		fmt.Println(line)
	}
}

func rotate(matrix [][]int) {
	last := len(matrix) - 1
	first := 0
	for first < last {
		for i := 0; first+i < last; i++ {
			t := matrix[first+i][last]
			matrix[first+i][last] = matrix[first][first+i]
			t, matrix[last][last-i] = matrix[last][last-i], t
			t, matrix[last-i][first] = matrix[last-i][first], t
			matrix[first][first+i] = t
		}
		first++
		last--
	}
}

func rotateKek(matrix [][]int) {
	transpose(matrix)
	reflect(matrix)
}

func transpose(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[j][i], matrix[i][j] = matrix[i][j], matrix[j][i]
		}
	}
}

func reflect(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}
