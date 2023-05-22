package main

import (
	"fmt"
)

func main() {
	fmt.Println(insert([][]int{{3, 5}, {12, 15}}, []int{6, 6}))
	// fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) < 1 {
		return [][]int{newInterval}
	}
	first, last := newInterval[0], newInterval[1]
	fi, li := 0, 0
	i := 0
	if first < intervals[0][0] && last < intervals[0][0] {
		return append([][]int{newInterval}, intervals...)
	}
	for ; i < len(intervals) && first >= intervals[0][0]; i++ {
		in := intervals[i]
		f, l := in[0], in[1]
		if first > l {
			continue
		} else if first >= f && first <= l {
			first = min(first, f)
			fi = i // врез по месту
			break
		} else {
			// first не трогаем
			fi = i // врез на 1 позже
			break
		}
	}
	for ; i < len(intervals); i++ {
		in := intervals[i]
		f, l := in[0], in[1]
		if f <= last && last <= l {
			last = max(last, l)
			li = i + 1
			break
		} else if last < l {
			li = i
			break
		} else {
			li = i + 1
		}
	}
	if fi == 0 && li == 0 {
		return append(intervals, newInterval)
	}
	a, b := make([][]int, fi), make([][]int, len(intervals)-li)
	copy(a, intervals[:fi])
	copy(b, intervals[li:])
	fin := append(a, []int{first, last})
	fin = append(fin, b...)
	return fin
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
