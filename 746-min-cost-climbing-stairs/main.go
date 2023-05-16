package main

import (
	"fmt"
)

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 29}))
}

var lenToCost []int

func minCostClimbingStairs(cost []int) int {
	lenToCost = make([]int, len(cost)+1)
	return min(rec(cost), rec(cost[1:]))
}

func rec(costs []int) int {
	if len(costs) <= 2 {
		return costs[0]
	}
	if lenToCost[len(costs)] != 0 {
		return lenToCost[len(costs)]
	}
	r1 := rec(costs[1:])
	r2 := rec(costs[2:])
	currentCost := costs[0] + min(r1, r2)
	lenToCost[len(costs)] = currentCost
	return currentCost
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
