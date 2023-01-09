package main

import (
	"fmt"
)

func main() {
	e1 := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(e1))
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if !isValid(board[i][j], i, j, board) {
				return false
			}
		}
	}
	return true
}

func isValid(num byte, line, column int, board [][]byte) bool {
	board[line][column] = 0
	// check line and column
	for i := 0; i < 9; i++ {
		if num == board[line][i] || num == board[i][column] {
			return false
		}
	}
	// check quadrant
	for i := line / 3 * 3; i < line/3*3+3; i++ {
		for j := column / 3 * 3; j < column/3*3+3; j++ {
			if num == board[i][j] {
				return false
			}
		}
	}
	board[line][column] = num
	return true
}
