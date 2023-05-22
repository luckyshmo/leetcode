package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{{1, 4}, {0, 2}, {3, 5}}))
	fmt.Println(merge([][]int{{1, 4}, {2, 3}}))
	fmt.Println(merge([][]int{{1, 4}, {5, 6}}))
	fmt.Println(merge([][]int{{0, 4}, {1, 4}}))
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}

func merge(in [][]int) [][]int {
	if len(in) < 2 {
		return in
	}
	sort.Slice(in, func(i, j int) bool {
		return in[i][0] < in[j][0]
	})

	ii := 1 // new array len
	for i := 0; i < len(in); i++ {
		if in[ii-1][1] >= in[i][0] {
			in[ii-1][1] = max(in[i][1], in[ii-1][1])
			continue
		}
		in[ii] = in[i]
		ii++
	}
	return in[:ii]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
