package main

import "fmt"

func main() {
	fmt.Println(isBipartite([][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}))
	fmt.Println(isBipartite([][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}))
	fmt.Println(isBipartite([][]int{{1}, {0, 3}, {3}, {1, 2}}))
}

func isBipartite(graph [][]int) bool {
	colors := make([]int, len(graph))

	for i := 0; i < len(graph); i++ {
		if colors[i] == 0 && !isValid(graph, colors, i, 1) {
			return false
		}
	}

	return true
}

func isValid(graph [][]int, colors []int, head int, color int) bool {
	if colors[head] != 0 {
		return colors[head] == color
	}

	colors[head] = color

	for _, currentNode := range graph[head] {
		if colors[head] == colors[currentNode] || !isValid(graph, colors, currentNode, -color) {
			return false
		}
	}
	return true
}
