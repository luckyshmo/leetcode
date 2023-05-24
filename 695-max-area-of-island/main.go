package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxAreaOfIsland([][]int{{1, 1, 1, 0, 0, 0, 1}}))
}

var area = 0

func maxAreaOfIsland(grid [][]int) (max int) {
	for c := 0; c < len(grid[0]); c++ {
		for r := 0; r < len(grid); r++ {
			if grid[r][c] == 1 {
				nullifyAndCount(grid, c, r)
				if area > max {
					max = area
				}
				area = 0
			}
		}
	}
	return
}

func nullifyAndCount(grid [][]int, c, r int) {
	if c < len(grid[0]) && c >= 0 &&
		r < len(grid) && r >= 0 &&
		grid[r][c] == 1 {
		area++
		grid[r][c] = 0
		nullifyAndCount(grid, c+1, r)
		nullifyAndCount(grid, c, r+1)
		nullifyAndCount(grid, c, r-1)
		nullifyAndCount(grid, c-1, r)
	}
}
