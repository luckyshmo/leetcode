package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isReflected([][]int{{0, 0}, {1, 0}, {3, 0}}))
}

func isReflected(points [][]int) bool {
	m := make(map[[2]int]struct{})
	var minXmaxY, maxXmaxY = []int{math.MaxInt, math.MinInt}, []int{math.MinInt, math.MinInt}
	for _, p := range points {
		maxXmaxY = MaxYmaxX(p, maxXmaxY)
		minXmaxY = MaxYminX(p, minXmaxY)
		m[[2]int{p[0], p[1]}] = struct{}{}
	}

	if !equalY(maxXmaxY, minXmaxY) {
		return false
	}

	symLine := float64(maxXmaxY[0]+minXmaxY[0]) / 2
	for _, p := range points {
		distToSymLine := mod(symLine - float64(p[0]))
		var p1 [2]int
		if float64(p[0]) >= symLine {
			p1 = [2]int{p[0] - int(distToSymLine*2), p[1]}
		} else {
			p1 = [2]int{p[0] + int(distToSymLine*2), p[1]}
		}
		if _, ok := m[p1]; !ok {
			return false
		}
	}

	return true
}

func equalY(p1, p2 []int) bool {
	return p1[1] == p2[1]
}

func MaxYmaxX(p1, p2 []int) []int {
	// if Y is equal, return by X
	if p1[1] == p2[1] {
		if p1[0] > p2[0] {
			return p1
		}
	} else if p1[1] > p2[1] {
		return p1
	}
	return p2
}

func MaxYminX(p1, p2 []int) []int {
	// if Y is equal, return by X
	if p1[1] == p2[1] {
		if p1[0] < p2[0] {
			return p1
		}
	} else if p1[1] > p2[1] {
		return p1
	}
	return p2
}

func mod(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}
