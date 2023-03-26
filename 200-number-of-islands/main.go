package main

import (
	"fmt"
)

func main() {
	fmt.Println(numIslands([][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}))
}

func numIslands(grid [][]byte) (res int) {
	for c := 0; c < len(grid[0]); c++ {
		for r := 0; r < len(grid); r++ {
			if grid[r][c] == '1' {
				nullify(grid, c, r)
				res++
			}
		}
	}
	return
}

func nullify(grid [][]byte, c, r int) {
	if c < len(grid[0]) && c >= 0 &&
		r < len(grid) && r >= 0 &&
		grid[r][c] == '1' {
		grid[r][c] = 0
		nullify(grid, c+1, r)
		nullify(grid, c, r+1)
		nullify(grid, c, r-1)
		nullify(grid, c-1, r)
	}
}
