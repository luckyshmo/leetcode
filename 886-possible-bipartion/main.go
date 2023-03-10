package main

import "fmt"

func main() {
	// fmt.Println(possibleBipartition(4, [][]int{{1, 2}, {1, 3}, {2, 4}}))
	// fmt.Println(possibleBipartition(3, [][]int{{1, 2}, {1, 3}, {2, 3}}))
	fmt.Println(possibleBipartition(10, [][]int{{1, 2}, {3, 4}, {5, 6}, {6, 7}, {8, 9}, {7, 8}}))
}

func possibleBipartition(n int, dislikes [][]int) bool {
	var graph = make([][]int, n)

	for _, pair := range dislikes {
		i := pair[0] - 1
		j := pair[1] - 1
		graph[i] = append(graph[i], j)
		graph[j] = append(graph[j], i)
	}
	var colors = make([]int, n)

	for i := 0; i < n; i++ {
		if colors[i] == 0 && !isValid(&graph, &colors, i, 1) {
			return false
		}
	}

	return true
}

func isValid(graph *[][]int, colors *[]int, i int, color int) bool {
	if (*colors)[i] != 0 {
		return (*colors)[i] == color
	}

	(*colors)[i] = color

	for _, j := range (*graph)[i] {
		if (*colors)[i] == (*colors)[j] || !isValid(graph, colors, j, -color) {
			return false
		}
	}

	return true
}
