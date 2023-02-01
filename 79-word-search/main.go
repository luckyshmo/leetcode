package main

import (
	"fmt"
)

func main() {
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCESEEEFS"))
	// fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE"))
	// fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))
}

func exist(board [][]byte, word string) (contains bool) {
	if word == "" {
		return true
	}

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if word[0] == board[y][x] {
				var visited [][]bool
				for i := range board {
					visited = append(visited, make([]bool, len(board[i])))
				}
				contains = search(board, word, x, y, visited)
				if contains {
					return
				}
			}
		}
	}
	return
}

func search(board [][]byte, word string, x, y int, visited [][]bool) bool {
	if x < 0 || y < 0 || y >= len(board) || x >= len(board[0]) {
		return false
	}
	if visited[y][x] {
		return false
	}
	if len(word) == 1 && word[0] == board[y][x] {
		return true
	} else if len(word) == 1 {
		return false
	} else if word[0] == board[y][x] {
		visited[y][x] = true
		defer func() {
			visited[y][x] = false
		}()
		if search(board, word[1:], x, y+1, visited) {
			return true
		}
		if search(board, word[1:], x, y-1, visited) {
			return true
		}
		if search(board, word[1:], x-1, y, visited) {
			return true
		}
		if search(board, word[1:], x+1, y, visited) {
			return true
		}
	}
	return false
}
