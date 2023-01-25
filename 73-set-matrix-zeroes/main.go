package main

import (
	"fmt"
)

func main() {
	m := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(m)
	fmt.Println(m)
}

func setZeroes(matrix [][]int) {
	colums, rows := make([]int, 0), make([]int, 0)
	for i, row := range matrix {
		for j, num := range row {
			if num == 0 {
				colums = append(colums, j)
				rows = append(rows, i)
			}
		}
	}

	for _, r := range rows {
		matrix[r] = make([]int, len(matrix[r]))
	}

	for _, column := range colums {
		for _, row := range matrix {
			row[column] = 0
		}
	}
}
